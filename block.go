package Hblocks

import (
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var keystr = `{"address":"2e11d0b3a9bd28cac42154b38b4b354f7fdb0738","crypto":{"cipher":"aes-128-ctr","ciphertext":"d587d1a3cc2f57a22191b3f41e9cde6d57b160847a18a6c06ebc6af5515db2f8","cipherparams":{"iv":"4da02a1efe9829a0c31b3cf419f2f89a"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"30b9e99fe9eba8265a5c50d53be393683ce4b8229ee6028e725713a94eae5eba"},"mac":"431f3ef0d0b3ec261359266b41828a92eccfe9ac201187cfda1508ecb9f7fc74"},"id":"aa5cd67e-0593-47b2-90b7-7cbb1bd092c1","version":3}`
var instance *Hello
var contractAddr = "0xa77F258b9Bd545d3a3dBB0803702Ec552afe34fB"

func init() {
	cli, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Panic("failed to Dial ", err)
	}

	hello, err := NewHello(common.HexToAddress(contractAddr), cli)
	if err != nil {
		log.Panic("Failed to NewHello ", err)
	}

	instance = hello
}

func getStr() (string, error) {
	return instance.GetStr(nil)
}

func setStr(str string) error {
	reader := strings.NewReader(keystr)
	auth, err := bind.NewTransactor(reader, "123")
	if err != nil {
		fmt.Println("Failed to NewTransactor ", err)
		return err
	}

	_, err = instance.SetStr(auth, str)
	if err != nil {
		fmt.Println("Failed to SetStr ", err)
		return err
	}
	return nil
}
