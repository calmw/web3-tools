/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"math/big"
	"os"
	"regexp"
)

// contractCallCmd represents the contractCall command
var contractCallCmd = &cobra.Command{
	Use:   "contractCall",
	Short: "合约调用",
	Long: `合约调用. 例子:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		abi, err := getContractAbiByContractName(contractName)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(args, 1111)

		function, err := getFuncInfoByFunctionName(contractName, functionName)
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(function.Para) > len(args) {
			fmt.Println("参数数量不匹配，参考：", function.Sig)
			return
		}

		parameter, err := parsePara(function.Para, args)

		fmt.Println(parameter, 678)

		fmt.Println(parameter, function.FuncName, 666)
		if err != nil {
			fmt.Printf("参数错误：%v \n", err)
			return
		}

		data, err := abi.Pack(function.FuncName, parameter...)
		if err != nil {
			fmt.Println("contract abi pack error:", err)
			return
		}

		fmt.Println(data)

		// 获取配置
		var configs = map[string]Config{}
		configNameBytes, err := FDB.Get([]byte("config_default"))
		if err != nil {
			fmt.Println("获取配置错误", err)
			return
		}
		configData, err := FDB.Get([]byte("config"))
		if err != nil {
			fmt.Println("获取配置错误", err)
			return
		}
		err = json.Unmarshal(configData, &configs)
		if err != nil {
			fmt.Println("提取配置错误", err)
			return
		}

		configDefault, ok := configs[string(configNameBytes)]
		if !ok {
			fmt.Println("提取默认配置错误")
			return
		}

		cli, err := ethclient.Dial(configDefault.RPC)
		if err != nil {
			fmt.Println("ethclient.Dial error : ", err)
			os.Exit(0)
		}

		tx, err := sendTx(cli, data, 88, contractAddress, configDefault.PrivateKey)
		if err != nil {
			fmt.Println("交易错误", err)
			return
		}
		marshalJSON, err := tx.MarshalJSON()
		if err != nil {
			fmt.Println("tx marshal 错误", err)
			return
		}

		fmt.Println(string(marshalJSON))
	},
}

func init() {
	rootCmd.AddCommand(contractCallCmd)

	contractCallCmd.Flags().StringVarP(&contractName, "contractName", "c", "", "合约名称")
	contractCallCmd.Flags().StringVarP(&contractAddress, "contractAddress", "a", "", "合约地址")
	contractCallCmd.Flags().StringVarP(&functionName, "functionName", "f", "", "合约方法名称")
	//contractCallCmd.Flags().StringVarP(&functionPara, "functionPara", "p", "", "方法参数")
	contractCallCmd.Flags().StringVarP(&ethNum, "ethNum", "v", "0", "ETH数量[可选参数]，payable方法需要次参数")
}

func genericTx(cli *ethclient.Client, contractAddress, wallet common.Address, data []byte, value int64) (*types.Transaction, error) {
	//获取nonce
	nonce, err := cli.NonceAt(context.Background(), wallet, nil)
	if err != nil {
		return nil, err
	}
	//获取小费
	gasTipCap, _ := cli.SuggestGasTipCap(context.Background())
	//transfer 默认是 使用 21000 gas
	gas := uint64(100000)
	//最大gas fee
	gasFeeCap := big.NewInt(38694000460)

	//创建交易
	tx := types.NewTx(&types.DynamicFeeTx{
		Nonce:     nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gas,
		To:        &contractAddress,
		Value:     big.NewInt(value),
		Data:      data,
	})
	return tx, nil
}

func sendTx(cli *ethclient.Client, data []byte, value int64, contractAddress, key string) (*types.Transaction, error) {

	wallet_, prv, err := getWalletFromKey(key)
	if err != nil {
		return nil, err
	}
	tx, err := genericTx(cli, common.HexToAddress(contractAddress), wallet_, data, value)
	if err != nil {
		return nil, err
	}

	// 获取当前区块链的ChainID
	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("获取ChainID失败:%v", err)
	}

	//创建签名者
	signer := types.NewLondonSigner(chainID)

	//对交易进行签名
	signTx, err := types.SignTx(tx, signer, prv)
	if err != nil {
		return nil, fmt.Errorf("签名失败:%v", err)
	}

	//发送交易
	err = cli.SendTransaction(context.Background(), signTx)
	if err != nil {
		return nil, fmt.Errorf("发送交易失败:%v", err)
	}

	return signTx, nil
}

func parsePara(abiPara, para []string) ([]interface{}, error) {
	var res []interface{}
	for i, ap := range abiPara {
		switch ap {
		case "address":
			if !isValidWalletAddress(para[i]) {
				return nil, fmt.Errorf("第%d个参数错误,地址格式错误", i+1)
			}
			res = append(res, common.HexToAddress(para[i]))
		case "uint256":
			bigint := new(big.Int)
			bigint.SetString(para[i], 10)
			res = append(res, bigint)
		}

	}

	return res, nil
}

func isValidWalletAddress(address string) bool {
	// 定义钱包地址的正则表达式规则
	pattern := "^0x[a-fA-F0-9]{40}$" // Ethereum钱包地址格式为0x开头后接40位十六进制字符串

	// 创建正则表达式对象并编译模式
	regExp := regexp.MustCompile(pattern)

	// 判断输入的地址是否与正则表达式匹配
	if regExp.MatchString(address) {
		return true
	} else {
		return false
	}
}

func getWalletFromKey(key string) (common.Address, *ecdsa.PrivateKey, error) {
	//从私钥推导出 公钥
	privateKeyECDSA, err := crypto.HexToECDSA(key)
	if err != nil {
		fmt.Println("crypto.HexToECDSA error ,", err)
		return common.Address{}, nil, fmt.Errorf("crypto.HexToECDSA error %v", err)
	}
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, nil, fmt.Errorf("publicKeyECDSA error %v", err)
	}
	//从公钥推导出钱包地址
	wallet_ := crypto.PubkeyToAddress(*publicKeyECDSA)

	return wallet_, privateKeyECDSA, nil
}

func getContractAbiByContractName(contractName string) (abi.ABI, error) {
	var contractsMap map[string]Contract
	contractData, err := FDB.Get([]byte("contracts"))

	err = json.Unmarshal(contractData, &contractsMap)
	if err != nil {
		return abi.ABI{}, err
	}

	//fmt.Println(contractsMap)
	contract, ok := contractsMap[contractName]
	if !ok {
		return abi.ABI{}, fmt.Errorf("该合约不存在")
	}

	contractAbi, err := abi.JSON(bytes.NewReader(contract.Abi))
	if err != nil {
		return abi.ABI{}, fmt.Errorf("abi.JSON error %v", err)
	}

	return contractAbi, nil
}

func getFuncInfoByFunctionName(contractName, functionName string) (ContractFunc, error) {
	var contractsMap map[string]Contract
	contractData, err := FDB.Get([]byte("contracts"))

	err = json.Unmarshal(contractData, &contractsMap)
	if err != nil {
		return ContractFunc{}, err
	}

	contract, ok := contractsMap[contractName]
	if !ok {
		return ContractFunc{}, fmt.Errorf("该合约不存在")
	}

	contractFunc, ok := contract.FuncMap[functionName]
	if !ok {
		return ContractFunc{}, fmt.Errorf("该方法不存在")
	}

	return contractFunc, nil
}
