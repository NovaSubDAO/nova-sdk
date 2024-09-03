package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CreateTransaction(from common.Address, to common.Address, data []byte, rpcEndpoint string) (*types.Transaction, error) {
	client, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		log.Fatalf("Failed to connect to the Optimism client: %v", err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), from)
	if err != nil {
		return nil, fmt.Errorf("Failed to get nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Failed to suggest gas price: %v", err)
	}

	msg := ethereum.CallMsg{From: from, To: &to, GasPrice: gasPrice, Value: big.NewInt(0), Data: data}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		log.Printf("Gas estimation failed, using fallback gas limit: %v", err)
		gasLimit = 20000000
	}

	tx := types.NewTransaction(nonce, to, big.NewInt(0), gasLimit, gasPrice, data)
	if err != nil {
		return nil, fmt.Errorf("Failed to create transaction: %v", err)
	}

	fmt.Println("Nonce: ", tx.Nonce())

	return tx, nil
}

func SignTransaction(tx *types.Transaction, chainId *big.Int, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		return nil, fmt.Errorf("Failed to sign transaction: %v", err)
	}

	return signedTx, nil
}

func SendTransaction(signedTx *types.Transaction, rpcEndpoint string) error {
	client, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		log.Fatalf("Failed to connect to the Optimism client: %v", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return fmt.Errorf("Failed to send transaction: %v", err)
	}

	return nil
}
