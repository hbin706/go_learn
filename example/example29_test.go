package example

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"testing"
)

func Test_29(t *testing.T) {

	//您可以使用模拟客户端来快速轻松地在本地测试您的交易，非常适合单元测试。为了开始，我们需要一个带有初始 ETH 的账户。为此，首先生成一个账户私钥。

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	//接着从 accounts/abi/bind 包创建一个 NewKeyedTransactor，并为其传递私钥。
	auth := bind.NewKeyedTransactor(privateKey)

	//下一步是创建一个创世账户并为其分配初始余额。我们将使用 core 包的 GenesisAccount 类型。
	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10) // 10 eth in wei

	address := auth.From
	genesisAlloc := map[common.Address]core.GenesisAccount{
		address: {
			Balance: balance,
		},
	}
	//现在我们将创世分配结构体和配置好的汽油上限传给 account/abi/bind/backends 包的 NewSimulatedBackend 方法，该方法将返回一个新的模拟以太坊客户端。
	blockGasLimit := uint64(4712388)
	client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)

	//您可以像往常一样使用此客户端。作为一个示例，我们将构造一个新的交易并进行广播。
	fromAddress := auth.From
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID := big.NewInt(1)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex()) // tx sent: 0xec3ceb05642c61d33fa6c951b54080d1953ac8227be81e7b5e4e2cfed69eeb51

	//到现在为止，您可能想知道交易何时才会被开采。为了“开采”它，您还必须做一件额外的事情，在客户端调用 Commit 提交新开采的区块。
	client.Commit()

	receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt == nil {
		log.Fatal("receipt is nil. Forgot to commit?")
	}

	fmt.Printf("status: %v\n", receipt.Status) // status: 1
}
