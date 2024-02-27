package service

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

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

func EthBalance(cli *ethclient.Client, wallet string) (*big.Int, error) {
	return cli.BalanceAt(context.Background(), common.HexToAddress(wallet), nil)
}
