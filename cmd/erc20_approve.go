/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/calmw/web3-tools/blockchain/service"
	"github.com/spf13/cobra"
	"strings"
)

// erc20ApproveCmd represents the erc20Approve command
var erc20ApproveCmd = &cobra.Command{
	Use:   "erc20Approve",
	Short: "Approve 授权",
	Long: `给某账户Approve授权. 例子:

web3-tools erc20Approve --key 88efa0c968693a034301450d450c7169b3f608966977a4e3b5b0672cb82c597e --chainId 698 --spender 0x20142400c97AD2c0A74d4C0400e482a973087033  --rpc https://lisbon-testnet-rpc.matchtest.co  --token 0x3eE243ff68074502b1D9D65443fa97b99f634570 --wallet 0xD5f92Fd92F8c7f9391513E3019D9441aAf5b2D9E --amount 100`,

	Run: func(cmd *cobra.Command, args []string) {
		err, Cli := service.Client(RPC)
		if err != nil {
			fmt.Printf("错误：%s\n", err)
			return
		}
		Erc20 := service.NewErc20(Cli, token)
		privateKey = strings.TrimLeft(privateKey, "0x")
		tx, err := Erc20.Approve(Cli, chainId, privateKey, spender, amount)
		if err != nil {
			fmt.Printf("错误：%s\n", err)
			return
		}
		fmt.Printf("%-25s%d\n", "Nonce:", tx.Nonce())
		fmt.Printf("%-25s%s\n", "Hash:", tx.Hash())
	},
}

func init() {
	rootCmd.AddCommand(erc20ApproveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// erc20ApproveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// erc20ApproveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	erc20ApproveCmd.PersistentFlags().StringVarP(&RPC, "rpc", "r", "", "节点RPC")
	erc20ApproveCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "ERC20 合约地址")
	//erc20ApproveCmd.PersistentFlags().StringVarP(&wallet, "wallet", "w", "", "钱包地址")
	erc20ApproveCmd.PersistentFlags().Int64VarP(&chainId, "chainId", "i", 0, "链的chain ID")
	erc20ApproveCmd.PersistentFlags().StringVarP(&privateKey, "key", "k", "", "钱包私钥")
	erc20ApproveCmd.PersistentFlags().StringVarP(&spender, "spender", "s", "", "被授权的合约地址")
	erc20ApproveCmd.PersistentFlags().Int64VarP(&amount, "amount", "a", 0, "授权的数量")
}
