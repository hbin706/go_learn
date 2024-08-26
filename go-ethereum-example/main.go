package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

func main() {

	/*
		1。连接客户端
	*/
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections

	/*
		2。连接账户
	*/
	address := common.HexToAddress("0xEe0E144FE7dE487C6b137689b3AA312411a0760C")

	hasher := sha256.New()
	hasher.Write(address[:])
	hash := hasher.Sum(nil)
	fmt.Printf("%x\n", hash)   // 输出哈希值的十六进制形式
	fmt.Println(address.Hex()) // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F
	fmt.Println(address.Bytes())

	/*
		3。账户余额
	*/

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // 25893180161173005034

	blockNumber := big.NewInt(5532993)
	balanceAt, err := client.BalanceAt(context.Background(), address, blockNumber)
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println(balanceAt) // 25 729324269165216042

	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041

	pendingBalance, err := client.PendingBalanceAt(context.Background(), address)
	fmt.Println(pendingBalance) // 25729324269165216042
}
