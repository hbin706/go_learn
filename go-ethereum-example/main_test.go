package main

//
//import (
//	"context"
//	"crypto/ecdsa"
//	"fmt"
//	"github.com/ethereum/go-ethereum/accounts/abi"
//	"log"
//	"math/big"
//	"strings"
//	"testing"
//	"time"
//
//	"go_learn/store"
//
//	"github.com/ethereum/go-ethereum"
//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/core/types"
//	"github.com/ethereum/go-ethereum/crypto"
//	"github.com/ethereum/go-ethereum/ethclient"
//)
//
//func TestQueryBlock(t *testing.T) {
//	devClient, err := ethclient.Dial("http://localhost:7545")
//	if err != nil {
//		t.Fatal(err)
//	}
//	ctx := context.Background()
//	context.WithTimeout(ctx, 30*time.Second)
//	genesisBlock, err := devClient.BlockByNumber(ctx, big.NewInt(0))
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Logf("Genesis Block: %v\n", genesisBlock)
//}
//
//func TestDeployContract(t *testing.T) {
//	client, err := ethclient.Dial("http://localhost:7545")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	privateKey, err := crypto.HexToECDSA("d87c2f94c1c91a3aa90c4a70930bcafb5f8925310d77de4dd90d1fdceefd1bd4")
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
//	address, tx, instance, err := store.DeployStore(auth, client, input)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(address.Hex())
//	fmt.Println(tx.Hash().Hex())
//
//	_ = instance
//}
//
//func TestSetItem(t *testing.T) {
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
//		log.Fatal(err)
//	}
//	auth.Nonce = big.NewInt(int64(nonce))
//	auth.Value = big.NewInt(0)     // in wei
//	auth.GasLimit = uint64(300000) // in units
//	auth.GasPrice = gasPrice
//
//	address := common.HexToAddress("0x5fbdb2315678afecb367f032d93f642f64180aa3")
//	instance, err := store.NewStore(address, client)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	key := [32]byte{}
//	value := [32]byte{}
//	copy(key[:], []byte("foo"))
//	copy(value[:], []byte("bar"))
//
//	tx, err := instance.SetItem(auth, key, value)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("tx sent: %s", tx.Hash().Hex())
//
//	result, err := instance.Items(&bind.CallOpts{From: fromAddress}, key)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("-----", string(result[:]))
//}
//
//func TestGetItem(t *testing.T) {
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
//		log.Fatal(err)
//	}
//	auth.Nonce = big.NewInt(int64(nonce))
//	auth.Value = big.NewInt(0)     // in wei
//	auth.GasLimit = uint64(300000) // in units
//	auth.GasPrice = gasPrice
//
//	address := common.HexToAddress("0x5fbdb2315678afecb367f032d93f642f64180aa3")
//	instance, err := store.NewStore(address, client)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	key := [32]byte{}
//	copy(key[:], []byte("foo"))
//
//	tx, err := instance.GetItem(auth, key)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
//
//	result, err := instance.Items(&bind.CallOpts{From: fromAddress}, key)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("-----", string(result[:]))
//}
//
//func TestQueryEvent(t *testing.T) {
//	client, err := ethclient.Dial("http://localhost:8545")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	contractAddress := common.HexToAddress("0x5fbdb2315678afecb367f032d93f642f64180aa3")
//	query := ethereum.FilterQuery{
//		FromBlock: big.NewInt(1),
//		ToBlock:   big.NewInt(15),
//		Addresses: []common.Address{
//			contractAddress,
//		},
//		Topics: [][]common.Hash{
//			//{common.BytesToHash(crypto.Keccak256([]byte("ItemSet(bytes32,bytes32)")))},
//			{common.BytesToHash(crypto.Keccak256([]byte("ItemGet(bytes32)")))},
//		},
//	}
//
//	logs, err := client.FilterLogs(context.Background(), query)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	contractAbi, err := abi.JSON(strings.NewReader(store.StoreMetaData.ABI))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, vLog := range logs {
//		fmt.Println(vLog.BlockHash.Hex())
//		fmt.Println(vLog.BlockNumber)
//		fmt.Println(vLog.TxHash.Hex())
//
//		//event := struct {
//		//	Key   [32]byte
//		//	Value [32]byte
//		//}{}
//		//var event store.StoreItemSet
//		var event store.StoreItemGet
//		//err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
//		err := contractAbi.UnpackIntoInterface(&event, "ItemGet", vLog.Data)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		//fmt.Println(string(event.Key[:]))   // foo
//		//fmt.Println(string(event.Value[:])) // bar
//
//		fmt.Println(string(event.Key[:])) // foo
//		//fmt.Println(string(event.Value[:])) // bar
//
//		var topics [4]string
//		for i := range vLog.Topics {
//			topics[i] = vLog.Topics[i].Hex()
//		}
//
//		fmt.Println("------------topics=", topics[0])
//	}
//
//	//eventSignature := []byte("ItemSet(bytes32,bytes32)")
//	eventSignature := []byte("ItemGet(bytes32)")
//	hash := crypto.Keccak256Hash(eventSignature)
//	fmt.Println("---------eventSign=", hash.Hex())
//}
//
//func TestSubscribeLatestBlock(t *testing.T) {
//	client, err := ethclient.Dial("ws://localhost:8545")
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	headers := make(chan *types.Header)
//	sub, err := client.SubscribeNewHead(context.Background(), headers)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	for {
//		select {
//		case err := <-sub.Err():
//			log.Fatal(err)
//		case header := <-headers:
//			fmt.Println(header.Hash().Hex())
//
//			block, err := client.BlockByHash(context.Background(), header.Hash())
//			if err != nil {
//				log.Fatal(err)
//			}
//
//			fmt.Println(block.Hash().Hex())
//			fmt.Println(block.Number().Uint64())
//			fmt.Println(block.Time())
//			fmt.Println(block.Nonce())
//			fmt.Println(len(block.Transactions()))
//		}
//	}
//}
//
//func TestTransfer(t *testing.T) {
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
//	value := big.NewInt(1000000000000000000) // in wei (1 eth)
//	gasLimit := uint64(21000)                // in units
//	gasPrice, err := client.SuggestGasPrice(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	toAddress := common.HexToAddress("0x54706F2c125d1Bb5c67Ec91621d145F014d5E66e")
//	var data []byte
//	tx := types.NewTx(
//		&types.LegacyTx{
//			Nonce:    nonce,
//			To:       &toAddress,
//			Value:    value,
//			Gas:      gasLimit,
//			GasPrice: gasPrice,
//			Data:     data,
//		},
//	)
//
//	chainID, err := client.NetworkID(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = client.SendTransaction(context.Background(), signedTx)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
//}
//
//func TestSubscribeEvent(t *testing.T) {
//	client, err := ethclient.Dial("ws://localhost:8545")
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
//	queryGet := ethereum.FilterQuery{
//		Addresses: []common.Address{contractAddress},
//		Topics: [][]common.Hash{
//			{
//				common.BytesToHash(crypto.Keccak256([]byte("ItemGet(bytes32)"))),
//			},
//		},
//	}
//
//	querySet := ethereum.FilterQuery{
//		Addresses: []common.Address{contractAddress},
//		Topics: [][]common.Hash{
//			{
//				common.BytesToHash(crypto.Keccak256([]byte("ItemSet(bytes32,bytes32)"))),
//			},
//		},
//	}
//
//	getLogs := make(chan types.Log)
//	ctx := context.Background()
//	subGet, err := client.SubscribeFilterLogs(ctx, queryGet, getLogs)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	setLogs := make(chan types.Log)
//	subSet, err := client.SubscribeFilterLogs(ctx, querySet, setLogs)
//
//	for {
//		select {
//		case getErr := <-subGet.Err():
//			t.Fatal(getErr)
//		case setErr := <-subSet.Err():
//			t.Fatal(setErr)
//		case gLog := <-getLogs:
//			// 读取Item事件
//			fmt.Println(gLog)
//		case sLog := <-setLogs:
//			// 写入Item事件
//			fmt.Println(sLog)
//		case <-ctx.Done():
//			subGet.Unsubscribe()
//			subSet.Unsubscribe()
//			err = ctx.Err()
//			if err != nil {
//				t.Fatal(err)
//			}
//			return
//		}
//	}
//
//}
