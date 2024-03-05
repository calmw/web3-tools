/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// listConfigCmd represents the listConfig command
var listConfigCmd = &cobra.Command{
	Use:   "listConfig",
	Short: "列出配置信息",
	Long: `列出配置信息. 例子:

web3-tools listConfig`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := FDB.Get([]byte("config"))
		if err != nil {
			fmt.Println("获取配置错误", err)
			return
		}
		fmt.Println(string(data))

	},
}

func init() {
	rootCmd.AddCommand(listConfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listConfigCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listConfigCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
