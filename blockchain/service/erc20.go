package service

import (
	"context"
	erc20 "github.com/calmw/web3-tools/blockchain/bindding/token"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"
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

func Client(Rpc string) (error, *ethclient.Client) {
	client, err := ethclient.Dial(Rpc)
	if err != nil {
		return err, nil
	}
	return nil, client
}

func GetAuth(cli *ethclient.Client, priKey string, chainId int64) (*bind.TransactOpts, error) {
	privateKeyEcdsa, err := crypto.HexToECDSA(priKey)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyEcdsa, big.NewInt(chainId))
	if err != nil {
		return nil, err
	}

	return &bind.TransactOpts{
		From:      auth.From,
		Nonce:     nil,
		Signer:    auth.Signer,
		Value:     big.NewInt(0),
		GasPrice:  nil,
		GasFeeCap: nil,
		GasTipCap: nil,
		GasLimit:  0,
		Context:   context.Background(),
		NoSend:    false,
	}, nil
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
