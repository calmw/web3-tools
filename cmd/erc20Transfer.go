/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/calmw/web3-tools/blockchain/service"
	"strings"

	"github.com/spf13/cobra"
)

// erc20TransferCmd represents the erc20Transfer command
var erc20TransferCmd = &cobra.Command{
	Use:   "erc20Transfer",
	Short: "ERC20代币转账",
	Long: `ERC20代币转账. 例子:

 web3-tools erc20Transfer --rpc https://lisbon-testnet-rpc.matchtest.co --chainId 698 --token 0x3eE243ff68074502b1D9D65443fa97b99f634570 --key 0x88efa0c968693a034301450d450c7169b3f608966977a4e3b5b0672cb82c597e --to 0xD5f92Fd92F8c7f9391513E3019D9441aAf5b2D9E --transferAmount 1`,
	Run: func(cmd *cobra.Command, args []string) {
		err, Cli := service.Client(RPC)
		if err != nil {
			fmt.Printf("错误：%s\n", err)
			return
		}
		Erc20 := service.NewErc20(Cli, token)
		privateKey = strings.TrimLeft(privateKey, "0x")
		tx, err := Erc20.Transfer(Cli, chainId, privateKey, to, transferAmount)
		if err != nil {
			fmt.Printf("错误：%s\n", err)
			return
		}
		fmt.Printf("%-25s%d\n", "Nonce:", tx.Nonce())
		fmt.Printf("%-25s%s\n", "Hash:", tx.Hash())
	},
}

func init() {
	rootCmd.AddCommand(erc20TransferCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// erc20TransferCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// erc20TransferCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	erc20TransferCmd.PersistentFlags().StringVarP(&RPC, "rpc", "r", "", "节点RPC")
	erc20TransferCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "ERC20 合约地址")
	erc20TransferCmd.PersistentFlags().Int64VarP(&chainId, "chainId", "i", 0, "链的chain ID")
	erc20TransferCmd.PersistentFlags().StringVarP(&privateKey, "key", "k", "", "钱包私钥")
	erc20TransferCmd.PersistentFlags().StringVarP(&to, "to", "s", "", "收币账户")
	erc20TransferCmd.PersistentFlags().StringVarP(&transferAmount, "transferAmount", "a", "0", "转账的数量")

}
