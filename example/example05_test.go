package example

//
//import (
//	"fmt"
//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/ethclient"
//	"log"
//	"math"
//	"math/big"
//	"testing"
//
//	token "./contracts_erc20" // for demo
//)
//
//func Test_05(t *testing.T) {
//	client, err := ethclient.Dial("http://127.0.0.1:8545")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Golem (GNT) Address
//	tokenAddress := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
//	instance, err := token.NewToken(tokenAddress, client)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	address := common.HexToAddress("0x0536806df512d6cdde913cf95c9886f65b1d3462")
//	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
//	if err != nil {
//		log.Fatal(err)
//	}
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
//	fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"
//
//	fbal := new(big.Float)
//	fbal.SetString(bal.String())
//	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
//
//	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
//}
