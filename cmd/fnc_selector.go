// Package cmd
/*
Copyright © 2024 NAME HERE <calm.wang@hotmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
	"github.com/status-im/keycard-go/hexutils"
	"strings"
)

// getFnSelectorCmd represents the getFnSelector command
var getFnSelectorCmd = &cobra.Command{
	Use:   "getFnSelector",
	Short: "获取函数选择器,返回16进制字符串",
	Long: `根据函数签名获取函数选择器16进制值. 例子:
web3-tools getFnSelector --funcSig adminSetClaimType(uint256,uint256)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%-26s%s\n", "函数签名:", funcSig)
		fmt.Printf("%-25s%s\n", "函数选择器:", "0x"+strings.ToLower(hexutils.BytesToHex(crypto.Keccak256([]byte(funcSig))[:4])))
	},
}

func init() {
	rootCmd.AddCommand(getFnSelectorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getFnSelectorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getFnSelectorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	getFnSelectorCmd.PersistentFlags().StringVarP(&funcSig, "funcSig", "s", "", "方法签名")
}
