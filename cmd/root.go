// Package cmd
/*
Copyright © 2024 calm.wang@hotmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/calmw/fdb"
	"os"

	"github.com/spf13/cobra"
)

type Config struct {
	Name       string
	RPC        string
	ChainID    int64
	PrivateKey string
}

// var RPC, token, funcSig, someStr, wallet, owner, spender, privateKey, to, transferAmount, configName, contractName, contractAddress, functionName, functionPara string
var RPC, token, funcSig, someStr, wallet, owner, spender, privateKey, to, transferAmount, configName, contractName, contractAddress, functionName, ethNum string
var chainId, amount int64
var FDB *fdb.DB

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "web3-tools",
	Short: "一个web3 辅助工具",
	Long: `web3开发者的辅助工具. 例如:

生成函数选择器
对字符串进行keccak256加密
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.eth-tools.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	opts := fdb.DefaultOption
	opts.DirPath = "./web3-tools/fdb"
	db, err := fdb.Open(opts)
	if err != nil {
		panic(err)
	}
	FDB = db
}
