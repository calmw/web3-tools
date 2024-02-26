/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/calmw/web3-tools/blockchain/service"
	"math"
	"math/big"

	"github.com/spf13/cobra"
)

// erc20AllowanceCmd represents the erc20Allowance command
var erc20AllowanceCmd = &cobra.Command{
	Use:   "erc20Allowance",
	Short: "Erc20 查看Approve授权额度",
	Long: `Erc20 查看Approve授权额度. 例子:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err, Cli := service.Client(RPC)
		if err != nil {
			fmt.Printf("错误：%s\n", err)
			return
		}
		Erc20 := service.NewErc20(Cli, token)
		allowance, err := Erc20.Allowance(owner, spender)
		if err != nil {
			fmt.Printf("错误：%s\n", err)
			return
		}
		decimals, err := Erc20.Decimals()
		if err != nil {
			fmt.Printf("错误：%s\n", err)
			return
		}
		decimal := int64(math.Pow(10, float64(decimals)))
		fmt.Printf("%-26s%s\n", "Allowance:", allowance.String())
		fmt.Printf("%d%-20s%s\n", decimals, "位精度值:", allowance.Div(allowance, big.NewInt(decimal)).String())
	},
}

func init() {
	rootCmd.AddCommand(erc20AllowanceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// erc20AllowanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// erc20AllowanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	erc20AllowanceCmd.PersistentFlags().StringVarP(&RPC, "rpc", "r", "", "节点RPC")
	erc20AllowanceCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "ERC20 合约地址")
	erc20AllowanceCmd.PersistentFlags().StringVarP(&owner, "owner", "o", "", "钱包地址")
	erc20AllowanceCmd.PersistentFlags().StringVarP(&spender, "spender", "s", "", "被授权的合约地址")
}

// go run web3-tools.go erc20Allowance --rpc https://lisbon-testnet-rpc.matchtest.co  --token 0x3eE243ff68074502b1D9D65443fa97b99f634570 --owner 0xD5f92Fd92F8c7f9391513E3019D9441aAf5b2D9E --spender 0x20142400c97AD2c0A74d4C0400e482a973087033
