package main

import (
	"log"
	"math/big"

	"github.com/NovaSubDAO/nova-sdk/go/pkg/sdk"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// check vault address on ethereum
	log.Println("Nova SDK: Ethereum")
	chainId := int64(1)
	novaSdk, err := sdk.NewNovaSDK("", chainId)
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}
	log.Println("chain id:", novaSdk.Config.ChainId)
	log.Println("vault address:", novaSdk.Config.VaultAddress)
	amount := big.NewInt(100)
	referral := big.NewInt(123)
	tx, err := novaSdk.SdkDomain.CreateDepositTransaction(common.HexToAddress("0x83F20F44975D03b1b09e64809B757c47f942BEeA"), common.HexToAddress("0x83F20F44975D03b1b09e64809B757c47f942BEeA"), amount, referral)
	if err != nil {
		log.Fatal("Error creating deposit transaction:", err)
	}
	log.Println("deposit tx:", tx)
	tx, err = novaSdk.SdkDomain.CreateWithdrawTransaction(common.HexToAddress("0x83F20F44975D03b1b09e64809B757c47f942BEeA"), common.HexToAddress("0x83F20F44975D03b1b09e64809B757c47f942BEeA"), amount, referral)
	if err != nil {
		log.Fatal("Error creating withdraw transaction:", err)
	}
	log.Println("withdraw tx:", tx)
	log.Println(" ")

	// check vault address on optimism
	log.Println("Nova SDK: Optimism")
	chainId = 10
	novaSdk, err = sdk.NewNovaSDK("", chainId)
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}
	log.Println("chain id:", novaSdk.Config.ChainId)
	log.Println("vault address:", novaSdk.Config.VaultAddress)
	tx, err = novaSdk.SdkDomain.CreateDepositTransaction(common.HexToAddress("0x83F20F44975D03b1b09e64809B757c47f942BEeA"), common.HexToAddress("0x83F20F44975D03b1b09e64809B757c47f942BEeA"), amount, referral)
	if err != nil {
		log.Fatal("Error creating deposit transaction:", err)
	}
	log.Println("deposit tx:", tx)
	tx, err = novaSdk.SdkDomain.CreateWithdrawTransaction(common.HexToAddress("0x83F20F44975D03b1b09e64809B757c47f942BEeA"), common.HexToAddress("0x83F20F44975D03b1b09e64809B757c47f942BEeA"), amount, referral)
	if err != nil {
		log.Fatal("Error creating withdraw transaction:", err)
	}
	log.Println("withdraw tx:", tx)
}
