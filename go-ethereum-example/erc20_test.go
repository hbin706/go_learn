package main

import (
	"testing"
)

func deploy(t *testing.T) {
	//client, err := ethclient.Dial("http://localhost:8545")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//contract, err := store.bindStore(
	//	common.HexToAddress("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"),
	//	caller, nil, nil)
}

//func TestDeployContract1(t *testing.T) {
//	client, err := ethclient.Dial("http://localhost:8545")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	publicKey := privateKey.Public()
//	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
//	if !ok {
//		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
//	}
//
//	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
//	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	gasPrice, err := client.SuggestGasPrice(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))
//	if err != nil {
//		t.Fatal(err)
//	}
//	auth.Nonce = big.NewInt(int64(nonce))
//	auth.Value = big.NewInt(0)      // in wei
//	auth.GasLimit = uint64(3000000) // in units
//	auth.GasPrice = gasPrice
//
//	input := "1.0"
//	address, tx, instance, err := store.(auth, client, input)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(address.Hex())
//	fmt.Println(tx.Hash().Hex())
//
//	_ = instance
//}

// 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
// 0xf4058a5ad1aef7c38423acb98b14bf3dd4c597299f151ce4e7c15d75a9828948
//func TestItem1(t *testing.T) {
//
//	client, err := ethclient.Dial("http://localhost:8545")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Golem (GNT) Address
//	tokenAddress := common.HexToAddress("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
//	instance, err := token.NewToken(tokenAddress, client)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	//address := common.HexToAddress("0xd238ff944bacb478cbed5efcae784d7bf4f2ff80")
//	//bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//
//	name, err := instance.Name(&bind.CallOpts{})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	symbol, err := instance.Symbol(&bind.CallOpts{})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	decimals, err := instance.Decimals(&bind.CallOpts{})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
//	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
//	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
//
//	//fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"
//
//	fbal := new(big.Float)
//	//fbal.SetString(bal.String())
//	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
//
//	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
//}
