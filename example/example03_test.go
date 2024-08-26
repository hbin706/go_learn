package example

import (
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func Test_03(t *testing.T) {
	address := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	fmt.Println(address.Hex()) // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F
	hasher := sha256.New()
	hasher.Write(address.Bytes())
	hash := hasher.Sum(nil)
	fmt.Println(hash)            // 0x00000000000000000000000071c7656ec7ab88b098defb751b7401b5f6d8976f
	fmt.Println(address.Bytes()) // [113 199 101 110 199 171 136 176 152 222 251 117 27 116 1 181 246 216 151 111]
}
