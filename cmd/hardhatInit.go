/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/calmw/fdb"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
	"github.com/status-im/keycard-go/hexutils"
	"os"
	"regexp"
	"strings"
)

type Hardhat struct {
	ProjectName  string
	ContractName []string // contracts目录下所有合约的名称
}

type HardHatABI struct {
	Format       string `json:"_format"`
	ContractName string `json:"contractName"`
	SourceName   string `json:"sourceName"`
	Abi          []struct {
		Inputs []struct {
			InternalType string `json:"internalType"`
			Name         string `json:"name"`
			Type         string `json:"type"`
		} `json:"inputs"`
		Name            string        `json:"name"`
		Type            string        `json:"type"`
		Anonymous       bool          `json:"anonymous,omitempty"`
		Outputs         []interface{} `json:"outputs,omitempty"`
		StateMutability string        `json:"stateMutability,omitempty"`
	} `json:"abi"`
	Bytecode         string `json:"bytecode"`
	DeployedBytecode string `json:"deployedBytecode"`
	LinkReferences   struct {
	} `json:"linkReferences"`
	DeployedLinkReferences struct {
	} `json:"deployedLinkReferences"`
}

type Package struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Main        string `json:"main"`
	Scripts     struct {
		Prettier string `json:"prettier"`
		Build    string `json:"build"`
		Debug    string `json:"debug"`
		Nodetest string `json:"nodetest"`
		Test     string `json:"test"`
	} `json:"scripts"`
	Keywords        []interface{} `json:"keywords"`
	Author          string        `json:"author"`
	License         string        `json:"license"`
	DevDependencies struct {
		NomiclabsHardhatEthers        string `json:"@nomiclabs/hardhat-ethers"`
		NomiclabsHardhatWaffle        string `json:"@nomiclabs/hardhat-waffle"`
		OpenzeppelinHardhatUpgrades   string `json:"@openzeppelin/hardhat-upgrades"`
		Chai                          string `json:"chai"`
		EthereumWaffle                string `json:"ethereum-waffle"`
		Ethers                        string `json:"ethers"`
		NomicfoundationHardhatToolbox string `json:"@nomicfoundation/hardhat-toolbox"`
		Hardhat                       string `json:"hardhat"`
		Prettier                      string `json:"prettier"`
		PrettierPluginSolidity        string `json:"prettier-plugin-solidity"`
	} `json:"devDependencies"`
	Dependencies struct {
		OpenzeppelinContracts            string `json:"@openzeppelin/contracts"`
		OpenzeppelinContractsUpgradeable string `json:"@openzeppelin/contracts-upgradeable"`
	} `json:"dependencies"`
}

type Contract struct {
	ContractName    string
	Abi             []byte
	FuncMap         map[string]ContractFunc // function name=>ContractFunc
	SelectorFuncMap map[string]ContractFunc // selector=>ContractFunc
}

type ContractFunc struct {
	FuncName   string
	Sig        string
	Para       []string
	FuncReturn []interface{}
	Selector   string
}

func newHardhat() *Hardhat {
	return &Hardhat{
		ContractName: make([]string, 0),
	}
}

// hardhatInitCmd represents the hardhatInit command
var hardhatInitCmd = &cobra.Command{
	Use:   "hardhatInit",
	Short: "Hardhat 项目帮助工具初始化",
	Long: `请切换到hardhat项目根目录，并在合约编译完成后执行. 例子:

web3-tools hardhatInit`,
	Run: func(cmd *cobra.Command, args []string) {
		///初始化DB
		opts := fdb.DefaultOption
		opts.DirPath = "./web3-tools/fdb"
		err := os.RemoveAll("./web3-tools")
		if err != nil {
			panic(err)
		}
		db, err := fdb.Open(opts)
		if err != nil {
			panic(err)
		}
		FDB = db
		// 判断当前是否在hardhat根目录
		_, err = os.Create("./web3-tools/contract.json")
		if err != nil {
			panic(err)
		}

		pwd, _ := os.Getwd()
		var isContractsDirExist bool
		var isArtifactsDirExist bool
		var isPackageJsonFilsExist bool

		dirEntries, err := os.ReadDir(pwd)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, entry := range dirEntries {
			if entry.IsDir() && entry.Name() == "contracts" {
				isContractsDirExist = true
			}
			if entry.IsDir() && entry.Name() == "artifacts" {
				isArtifactsDirExist = true
			}
			if !entry.IsDir() && entry.Name() == "package.json" {
				isPackageJsonFilsExist = true
			}
		}

		if !isContractsDirExist || !isPackageJsonFilsExist {
			fmt.Println("确定在hardhat根目录吗？没有的话，请在根目录下执行")
		}
		if !isArtifactsDirExist {
			fmt.Println("执行之前，请先试用hardhat编译合约")
		}

		// 创建目录
		err = os.MkdirAll("./web3-tools/abi", os.ModePerm)
		if err != nil {
			fmt.Println("创建ABI文件夹出错，error:", err)
			return
		}
		err = os.MkdirAll("./web3-tools/bin", os.ModePerm)
		if err != nil {
			fmt.Println("创建BIN文件夹出错，error:", err)
			return
		}

		//
		hardhat := newHardhat()
		err = hardhat.getProjectName()
		if err != nil {
			fmt.Println("获取项目信息出错，error:", err)
			return
		}

		err = hardhat.genericAbiAndBin()
		if err != nil {
			fmt.Println(err)
			return
		} // 生成ABI和BIN文件
		err = hardhat.genericGoBinding()
		if err != nil {
			fmt.Println(err)
			return
		} // 生成GO的绑定文件
		err = hardhat.genericContractDataAndSaveToDB()
		if err != nil {
			fmt.Println(err)
			return
		} // 将合约数据存入数据库

	},
}

func init() {
	rootCmd.AddCommand(hardhatInitCmd)

}

func (h *Hardhat) getProjectName() error {
	file, err := os.ReadFile("./package.json")
	if err != nil {
		return fmt.Errorf("读取package.json失败，error:%v", err)
	}
	var pkg Package
	err = json.Unmarshal(file, &pkg)
	if err != nil {
		return fmt.Errorf("解析package.json失败，error:%v", err)
	}

	h.ProjectName = formatPkgName(pkg.Name)

	return nil
}

func (h *Hardhat) genericAbiAndBin() error {
	dirEntries, err := os.ReadDir("./artifacts/contracts")
	if err != nil {
		return fmt.Errorf("读取artifacts目录下ABI目录失败，error:%v", err)
	}

	for _, entry := range dirEntries {
		entryName := entry.Name()
		if !strings.HasSuffix(entryName, ".sol") || !entry.IsDir() {
			continue
		}
		entries, err := os.ReadDir(fmt.Sprintf("./artifacts/contracts/%s", entryName))
		if err != nil {
			return fmt.Errorf("读取artifacts目录下ABI失败，error:%v", err)
		}
		for _, et := range entries {
			if et.IsDir() {
				continue
			}
			fileName := et.Name()
			if strings.HasSuffix(fileName, ".dbg.json") {
				continue
			}
			_, found := strings.CutSuffix(fileName, ".json")
			if !found {
				continue
			}
			// 读取hardhat生成的ABI Json文件
			abiFile := fmt.Sprintf("./artifacts/contracts/%s/%s", entryName, fileName)
			abiFileBytes, err := os.ReadFile(abiFile)
			if err != nil {
				return fmt.Errorf("读取ABI文件失败，error:%v", err)
			}
			var abi HardHatABI
			err = json.Unmarshal(abiFileBytes, &abi)
			if err != nil {
				return fmt.Errorf("解析ABI文件失败，error:%v", err)
			}
			// 添加合约名称
			h.ContractName = append(h.ContractName, abi.ContractName)
			// 生成新的ABI文件
			abiBytes, err := json.Marshal(abi.Abi)
			if err != nil {
				return fmt.Errorf("编码ABI文件失败，error:%v", err)
			}
			//生成json文件
			newAbiFile := fmt.Sprintf("./web3-tools/abi/%s.json", abi.ContractName)
			err = os.WriteFile(newAbiFile, abiBytes, os.ModePerm)
			if err != nil {
				return fmt.Errorf("生成ABI文件失败，error:%v", err)
			}
			// 生成新的BIN文件
			byteCodeBytes, err := json.Marshal(abi.Bytecode)
			if err != nil {
				return fmt.Errorf("编码BIN文件失败，error:%v", err)
			}
			//生成bin文件
			newBinFile := fmt.Sprintf("./web3-tools/bin/%s.bin", abi.ContractName)
			err = os.WriteFile(newBinFile, byteCodeBytes, os.ModePerm)
			if err != nil {
				return fmt.Errorf("生成BIN文件失败，error:%v", err)
			}
		}
	}

	return nil
}

func (h *Hardhat) genericGoBinding() error {
	dirEntries, err := os.ReadDir("./web3-tools/abi")
	if err != nil {
		return fmt.Errorf("读取web3-tools/abi目录下ABI文件失败，error:%v", err)
	}
	bindingDir := fmt.Sprintf("./web3-tools/binding/%s", h.ProjectName)
	err = os.MkdirAll(bindingDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("创建binding文件夹出错，error:%v", err)
	}

	for _, entry := range dirEntries {
		entryName := entry.Name()
		fileNameWithoutSuffix, found := strings.CutSuffix(entryName, ".json")
		if found {
			// 读取hardhat生成的ABI Json文件
			// abiFile, pkgName, typeName, outFileName string
			abiFileName := fmt.Sprintf("./web3-tools/abi/%s", entryName)
			binFileName := fmt.Sprintf("./web3-tools/bin/%s,bin", fileNameWithoutSuffix)
			typeName := strings.ToUpper(fileNameWithoutSuffix[:1]) + fileNameWithoutSuffix[1:]
			outFileName := fmt.Sprintf("./web3-tools/binding/%s/%s.go", h.ProjectName, fileNameWithoutSuffix)

			err = genericGoBinding(abiFileName, binFileName, h.ProjectName, typeName, outFileName)
			if err != nil {
				return fmt.Errorf("生成go绑定文件失败，error:%v", err)
			}
		} else {
			continue
		}
	}
	return nil
}

func (h *Hardhat) genericContractDataAndSaveToDB() error {
	var contractsMap = map[string]Contract{} // contract name => Contract
	dirEntries, err := os.ReadDir("./web3-tools/abi")
	if err != nil {
		return fmt.Errorf("读取web3-tools/abi目录下ABI文件失败，error:%v", err)
	}
	for _, entry := range dirEntries {
		abiBytes, err := os.ReadFile(fmt.Sprintf("./web3-tools/abi/%s", entry.Name()))
		if err != nil {
			return fmt.Errorf("读取web3-tools/abi目录下ABI文件失败，error:%v", err)
		}

		var abi HardHatABI
		err = json.Unmarshal(abiBytes, &abi.Abi)
		if err != nil {
			return fmt.Errorf("解码web3-tools/abi目录下ABI文件失败，error:%v", err)
		}
		contract := Contract{
			FuncMap:         map[string]ContractFunc{},
			SelectorFuncMap: map[string]ContractFunc{},
		}
		contractName_, _ := strings.CutSuffix(entry.Name(), ".json")
		contract.Abi = abiBytes
		contract.ContractName = contractName_
		for _, abi_ := range abi.Abi {
			if abi_.Type == "function" {
				contractFunc := ContractFunc{}
				contractFunc.FuncName = abi_.Name
				sig := abi_.Name + "("
				for _, input := range abi_.Inputs {
					sig += input.InternalType + ","
					contractFunc.Para = append(contractFunc.Para, input.InternalType)
				}
				sig = strings.TrimSuffix(sig, ",")
				sig += ")"
				fmt.Println(sig)
				contractFunc.Sig = sig

				for _, output := range abi_.Outputs {
					contractFunc.FuncReturn = append(contractFunc.FuncReturn, output)
				}

				contractFunc.Selector = "0x" + strings.ToLower(hexutils.BytesToHex(crypto.Keccak256([]byte(sig))[:4]))
				fmt.Println("0x" + strings.ToLower(hexutils.BytesToHex(crypto.Keccak256([]byte(sig))[:4])))
				contract.SelectorFuncMap[contractFunc.Selector] = contractFunc
				contract.FuncMap[contractFunc.FuncName] = contractFunc
			}
		}
		contractsMap[contract.ContractName] = contract
	}

	contractBytes, err := json.Marshal(contractsMap)
	if err != nil {
		return err
	}
	//
	err = FDB.Put([]byte("contracts"), contractBytes)
	if err != nil {
		return fmt.Errorf("内部错误：%v", err)
	}

	err = os.WriteFile("./web3-tools/contract.json", contractBytes, os.ModePerm)
	if err != nil {
		return fmt.Errorf("生成contract.json error：%v", err)
	}

	return nil
}

func genericGoBinding(abiFile, binFile, pkgName, typeName, outFileName string) error {
	var abis []string
	var bins []string
	var types []string
	var sigs []map[string]string
	var libs = make(map[string]string)
	var aliases = make(map[string]string)

	abi, err := os.ReadFile(abiFile)
	bin, err := os.ReadFile(binFile)
	abis = append(abis, string(abi))
	bins = append(bins, string(bin))
	types = append(types, typeName)
	code, err := bind.Bind(types, abis, bins, sigs, pkgName, bind.LangGo, libs, aliases)
	if err != nil {
		return fmt.Errorf("failed to generate HardHatABI binding: %v", err)
	}

	if err := os.WriteFile(outFileName, []byte(code), 0600); err != nil {
		return fmt.Errorf("failed to write HardHatABI binding: %v", err)
	}

	return nil
}

func formatPkgName(pkgName string) string {
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z]+`)
	return nonAlphanumericRegex.ReplaceAllString(pkgName, "")
}
