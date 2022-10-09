package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"
	"user-vote/contracts"
	"user-vote/dto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	DecimalOfPrecision = big.NewInt(1000000000000000000)
)

func Transfer(payment dto.Payment) bool {

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
		return false
	}

	privateKey, err := crypto.HexToECDSA(payment.KeySender)
	if err != nil {
		log.Fatal(err)
		return false
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return false
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
		return false
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return false
	}
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	addressR := common.HexToAddress(payment.Recipient)
	instance, err := contracts.NewContracts(addressR, client)
	valuePay := big.NewInt(payment.ValuePay)

	tx, err := instance.Transfer(auth, addressR, new(big.Int).Mul(valuePay, DecimalOfPrecision))
	if err != nil {
		log.Fatalf("Failed to transfer tokens, got error : %v", err)
		return false
	}
	fmt.Println("Created transaction ", tx)

	count, err := client.PendingTransactionCount(context.Background())
	if err != nil {
		log.Fatalf("Failed to get pending transactions, got error : %v", err)
		return false
	}
	fmt.Printf("Waiting for %d pending transactions \n", count)

	for {
		time.Sleep(3 * time.Second)
		fmt.Printf("Waiting receipt of transaction %s\n", tx.Hash().Hex())
		if !IsTransactionPending(client, context.Background(), tx.Hash()) {
			receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				log.Fatalf("Failed to get token info, got error : %v", err)
				return false
			}
			fmt.Printf("Got transaction receipt status: %v", receipt.Status)
			break
		}
	}
	return true
}

func IsTransactionPending(client *ethclient.Client, context context.Context, hash common.Hash) bool {
	_, pending, err := client.TransactionByHash(context, hash)
	if err != nil {
		panic(err)
	}
	return pending
}

func Balance(userAddress string) *big.Int {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(userAddress)
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	result := new(big.Int).Div(balance, DecimalOfPrecision)
	return result
}
