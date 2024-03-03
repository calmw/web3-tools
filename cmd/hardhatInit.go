/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/cobra"
	"os"
	"regexp"
	"strings"
)

type Hardhat struct {
	ProjectName  string
	ContractName []string // contracts目录下所有合约的名称
}

type ABI struct {
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
		// 判断当前是否在hardhat根目录
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

		hardhat.genericAbiAndBin()
		hardhat.genericAbiBind()
	},
}

func init() {
	rootCmd.AddCommand(hardhatInitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hardhatInitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hardhatInitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func (h *Hardhat) genericAbiAndBin() {
	dirEntries, err := os.ReadDir("./artifacts/contracts")
	if err != nil {
		fmt.Println("读取artifacts目录下ABI失败，error:", err)
		return
	}

	for _, entry := range dirEntries {
		entryName := entry.Name()
		fileNameWithoutSuffix, found := strings.CutSuffix(entryName, ".sol")
		if found {
			// 读取hardhat生成的ABI Json文件
			abiFile := fmt.Sprintf("./artifacts/contracts/%s/%s.json", entryName, fileNameWithoutSuffix)
			abiFileBytes, err := os.ReadFile(abiFile)
			if err != nil {
				fmt.Println("读取ABI文件失败，error:", err)
				return
			}
			var abi ABI
			err = json.Unmarshal(abiFileBytes, &abi)
			if err != nil {
				fmt.Println("解析ABI文件失败，error:", err)
				return
			}
			// 添加合约名称
			h.ContractName = append(h.ContractName, abi.ContractName)
			// 生成新的ABI文件
			abiBytes, err := json.Marshal(abi.Abi)
			if err != nil {
				fmt.Println("编码ABI文件失败，error:", err)
			}
			//生成json文件
			newAbiFile := fmt.Sprintf("./web3-tools/abi/%s.json", abi.ContractName)
			err = os.WriteFile(newAbiFile, abiBytes, os.ModePerm)
			if err != nil {
				fmt.Println("生成ABI文件失败，error:", err)
				return
			}
			// 生成新的BIN文件
			byteCodeBytes, err := json.Marshal(abi.Bytecode)
			if err != nil {
				fmt.Println("编码BIN文件失败，error:", err)
			}
			//生成bin文件
			newBinFile := fmt.Sprintf("./web3-tools/bin/%s.bin", abi.ContractName)
			err = os.WriteFile(newBinFile, byteCodeBytes, os.ModePerm)
			if err != nil {
				fmt.Println("生成BIN文件失败，error:", err)
				return
			}
		} else {
			continue
		}
	}
}

func (h *Hardhat) genericAbiBind() {
	dirEntries, err := os.ReadDir("./web3-tools/abi")
	if err != nil {
		fmt.Println("读取web3-tools/abi目录下ABI文件失败，error:", err)
		return
	}
	bindingDir := fmt.Sprintf("./web3-tools/binding/%s", h.ProjectName)
	err = os.MkdirAll(bindingDir, os.ModePerm)
	if err != nil {
		fmt.Println("创建binding文件夹出错，error:", err)
		return
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
				fmt.Println("生成go绑定文件失败，error:", err)
				return
			}

		} else {
			continue
		}
	}
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

func formatPkgName(pkgName string) string {
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z]+`)
	return nonAlphanumericRegex.ReplaceAllString(pkgName, "")
}

// Generate the contract binding
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
		return fmt.Errorf("failed to generate ABI binding: %v", err)
	}

	if err := os.WriteFile(outFileName, []byte(code), 0600); err != nil {
		return fmt.Errorf("failed to write ABI binding: %v", err)
	}

	return nil
}
