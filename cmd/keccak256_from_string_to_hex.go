/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/status-im/keycard-go/hexutils"
	"strings"

	"github.com/spf13/cobra"
)

// keccak256FromStringToHexCmd represents the keccak256FromStringToHex command
var keccak256FromStringToHexCmd = &cobra.Command{
	Use:   "keccak256FromStringToHex",
	Short: "对字符串进行keccak256加密，返回16进制字符串",
	Long: `对字符串进行keccak256加密，返回16进制字符串. 例子:
web3-tools keccak256FromStringToHex 'some string'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%-25s%s\n", "加密前:", args[0])
		fmt.Printf("%-25s%s\n", "加密后:", "0x"+strings.ToLower(hexutils.BytesToHex(crypto.Keccak256([]byte(args[0])))))
	},
}

func init() {
	rootCmd.AddCommand(keccak256FromStringToHexCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// keccak256FromStringToHexCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// keccak256FromStringToHexCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
