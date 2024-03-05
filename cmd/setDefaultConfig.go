/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// setDefaultConfigCmd represents the setDefaultConfig command
var setDefaultConfigCmd = &cobra.Command{
	Use:   "setDefaultConfig",
	Short: "将某个配置设置为默认配置",
	Long: `将某个配置设置为默认配置. 例子:

web3-tools setDefaultConfig match`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := SetDefaultConfig(configName); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("设置成功")
	},
}

func init() {
	rootCmd.AddCommand(setDefaultConfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setDefaultConfigCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setDefaultConfigCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	setDefaultConfigCmd.PersistentFlags().StringVarP(&configName, "configName", "n", "default", "配置名称")
}

func SetDefaultConfig(name string) error {
	var configs = map[string]Config{}

	// 获取配置
	data, err := FDB.Get([]byte("config"))
	if err != nil {
		return fmt.Errorf("获取现有配置错误%v", err)
	}
	err = json.Unmarshal(data, &configs)
	if err != nil {
		return fmt.Errorf("提取配置错误%v", err)
	}
	_, ok := configs[configName]
	if !ok {
		return errors.New("该配置不存在，请先设置")
	}

	err = FDB.Put([]byte("config_default"), []byte(configName))
	if err != nil {
		return fmt.Errorf("设置失败%v", err)
	}

	return nil
}
