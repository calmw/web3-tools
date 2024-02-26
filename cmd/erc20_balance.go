/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/calmw/web3-tools/blockchain/service"
	"github.com/spf13/cobra"
	"math"
	"math/big"
)

// erc20BalanceCmd represents the erc20Balance command
var erc20BalanceCmd = &cobra.Command{
	Use:   "erc20Balance",
	Short: "获取ERC20 token 余额",
	Long: `获取ERC20 token 余额. 例子:

web3-tools erc20Balance --rpc https://lisbon-testnet-rpc.matchtest.co  --token 0x3eE243ff68074502b1D9D65443fa97b99f634570 --wallet 0xD5f92Fd92F8c7f9391513E3019D9441aAf5b2D9E`,
	Run: func(cmd *cobra.Command, args []string) {
		err, Cli := service.Client(RPC)
		if err != nil {
			fmt.Printf("错误：%s\n", err)
			return
		}
		Erc20 := service.NewErc20(Cli, token)
		balance, err := Erc20.BalanceOf(wallet)
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
		fmt.Printf("%-25s%s\n", "余额:", balance.String())
		fmt.Printf("%d%-21s%s\n", decimals, "位精度值:", balance.Div(balance, big.NewInt(decimal)).String())
	},
}

func init() {
	rootCmd.AddCommand(erc20BalanceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// erc20BalanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// erc20BalanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	erc20BalanceCmd.PersistentFlags().StringVarP(&RPC, "rpc", "r", "", "节点RPC")
	erc20BalanceCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "ERC20 合约地址")
	erc20BalanceCmd.PersistentFlags().StringVarP(&wallet, "wallet", "w", "", "钱包地址")
}
