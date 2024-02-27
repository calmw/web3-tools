/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/calmw/web3-tools/blockchain/service"
	"math/big"

	"github.com/spf13/cobra"
)

// ethBalanceCmd represents the ethBalance command
var ethBalanceCmd = &cobra.Command{
	Use:   "ethBalance",
	Short: "获取ETH/native coin 余额",
	Long: `获取ETH/native coin 余额. 例子:

web3-tools ethBalance --rpc https://lisbon-testnet-rpc.matchtest.co  --wallet 0xD5f92Fd92F8c7f9391513E3019D9441aAf5b2D9E`,
	Run: func(cmd *cobra.Command, args []string) {
		err, Cli := service.Client(RPC)
		if err != nil {
			fmt.Printf("错误：%s\n", err)
			return
		}
		balance, err := service.EthBalance(Cli, wallet)
		if err != nil {
			fmt.Printf("错误：%s\n", err)
			return
		}

		fmt.Printf("%-25s%s\n", "余额:", balance.String())
		fmt.Printf("%-23s%s\n", "18位精度值:", balance.Div(balance, big.NewInt(1e18)).String())
	},
}

func init() {
	rootCmd.AddCommand(ethBalanceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ethBalanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ethBalanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	ethBalanceCmd.PersistentFlags().StringVarP(&RPC, "rpc", "r", "", "节点RPC")
	ethBalanceCmd.PersistentFlags().StringVarP(&wallet, "wallet", "w", "", "钱包地址")
}
