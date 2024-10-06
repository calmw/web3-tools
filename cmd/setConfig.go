/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// setConfigCmd represents the setConfig command
var setConfigCmd = &cobra.Command{
	Use:   "setConfig",
	Short: "配置设置",
	Long: `可以设置 chainID RPC privateKey 等全局配置，减少每次输入. 例子:

web3-tools setConfig match_test --rpc https://lisbon-testnet-rpc.matchtest.co --name match_test --chainId 698 --privateKey 80954a9*****`,
	Run: func(cmd *cobra.Command, args []string) {

		var configs = map[string]Config{}

		// 获取配置
		data, err := FDB.Get([]byte("config"))
		if err == nil {
			err = json.Unmarshal(data, &configs)
			if err != nil {
				fmt.Println("提取配置错误", err)
				return
			}
		}

		// 设置配置
		configs[configName] = Config{
			Name:       configName,
			RPC:        RPC,
			ChainID:    chainId,
			PrivateKey: privateKey,
		}

		marshal, err := json.MarshalIndent(configs, "", "\t")
		if err != nil {
			fmt.Println("数据错误", err)
			return
		}

		err = FDB.Put([]byte("config"), marshal)
		if err != nil {
			fmt.Println("配置错误", err)
			return
		}

		if err := SetDefaultConfig(configName); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("设置成功,当前默认配置设置为: %v \n", configName)
		fmt.Println("更改默认配置，请执行: setDefaultConfig --name \"config name\"")
	},
}

func init() {
	rootCmd.AddCommand(setConfigCmd)

	setConfigCmd.PersistentFlags().StringVarP(&configName, "name", "n", "default", "配置名称，比如可根据名称切换主网、测试网配置")
	setConfigCmd.PersistentFlags().StringVarP(&RPC, "rpc", "r", "https://lisbon-testnet-rpc.matchtest.co", "链RPC")
	setConfigCmd.PersistentFlags().Int64VarP(&chainId, "chainId", "i", 698, "链ID")
	setConfigCmd.PersistentFlags().StringVarP(&privateKey, "privateKey", "k", "", "私钥")
}
