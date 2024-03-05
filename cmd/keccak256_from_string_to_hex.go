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
web3-tools keccak256FromStringToHex --someStr "some string"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%-25s%s\n", "加密前:", someStr)
		fmt.Printf("%-25s%s\n", "加密后:", "0x"+strings.ToLower(hexutils.BytesToHex(crypto.Keccak256([]byte(someStr)))))
	},
}

func init() {
	rootCmd.AddCommand(keccak256FromStringToHexCmd)

	keccak256FromStringToHexCmd.PersistentFlags().StringVarP(&someStr, "someStr", "s", "", "字符串")
}
