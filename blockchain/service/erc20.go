package service

import (
	erc20 "github.com/calmw/web3-tools/blockchain/bindding/token"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"
	"strconv"
)

type Erc20 struct {
	Cli      *ethclient.Client
	Contract *erc20.Erc20
}

func NewErc20(cli *ethclient.Client, erc20Token string) *Erc20 {
	token, _ := erc20.NewErc20(common.HexToAddress(erc20Token), cli)

	return &Erc20{
		Cli:      cli,
		Contract: token,
	}
}

func (e Erc20) BalanceOf(wallet string) (*big.Int, error) {
	return e.Contract.BalanceOf(nil, common.HexToAddress(wallet))
}

func (e Erc20) Decimals() (uint8, error) {
	return e.Contract.Decimals(nil)
}

func (e Erc20) Approve(cli *ethclient.Client, chainId int64, priKey, spender string, amount int64) (*types.Transaction, error) {
	auth, err := GetAuth(cli, priKey, chainId)
	if err != nil {
		return nil, err
	}
	decimals, err := e.Decimals()
	if err != nil {
		return nil, err
	}

	decimal := int64(math.Pow(10, float64(decimals)))
	amountBigInt := big.NewInt(amount)
	amountBigInt.Mul(amountBigInt, big.NewInt(decimal))

	return e.Contract.Approve(auth, common.HexToAddress(spender), amountBigInt)
}

func (e Erc20) Allowance(owner, spender string) (*big.Int, error) {
	return e.Contract.Allowance(nil, common.HexToAddress(owner), common.HexToAddress(spender))
}

func (e Erc20) Transfer(cli *ethclient.Client, chainId int64, priKey, to string, amount string) (*types.Transaction, error) {
	auth, err := GetAuth(cli, priKey, chainId)
	if err != nil {
		return nil, err
	}

	f, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return nil, err
	}

	decimals, err := e.Decimals()
	if err != nil {
		return nil, err
	}

	amountBigInt := big.NewInt(int64(math.Pow(10, float64(decimals)) * f))

	return e.Contract.Transfer(auth, common.HexToAddress(to), amountBigInt)
}
