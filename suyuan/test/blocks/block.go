package blocks

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var keystr = `{"address":"34729d16e395ed5e8c8cdaa9fd89103261761bee","crypto":{"cipher":"aes-128-ctr","ciphertext":"0fe56c3daf35dc50a0093d11e6348a8e90eab23c2c2ee93b79daa0e0d7293213","cipherparams":{"iv":"21d7f2b3806575f0fe766acccc761c0f"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"012b9a1363e884ec3f24519fa3e03fa571393cc0cdc144fac7ac759a575917d2"},"mac":"2470b53830e88b60753e40d6ba2f74d50d94aa39044c4835c5936cdddeed6d1d"},"id":"f9b30b80-841b-46ac-9ea7-994eb83af838","version":3}`
var instance *Traceability
var contractAddr = "0x086435E06a4D1032F06B14D915467a4b4247DdD5"

// 初始化
func init() {
	cli, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Panic("failed to Dial ", err)
	}

	traceability, err := NewTraceability(common.HexToAddress(contractAddr), cli)
	if err != nil {
		log.Panic("Failed to NewHello ", err)
	}

	instance = traceability
}

func Register(username, passwd string) error {
	reader := strings.NewReader(keystr)
	auth, err := bind.NewTransactor(reader, "123")
	if err != nil {
		fmt.Println("failed to NewTransactor ", err)
		return err
	}
	_, err = instance.Register(auth, username, passwd)
	if err != nil {
		fmt.Println("failed to Register ", err)
		return err
	}
	return nil
}

func Login(username, passwd string) (bool, error) {
	return instance.Login(nil, username, passwd)
}

func Create_Goods(username, passwd, data, goodsID string) error {
	reader := strings.NewReader(keystr)
	auth, err := bind.NewTransactor(reader, "123")
	if err != nil {
		fmt.Println("failed to NewTransactor ", err)
		return err
	}
	_, err = instance.CreateGoods(auth, username, passwd, data, goodsID)
	if err != nil {
		fmt.Println("failed to CreateGoods ", err)
		return err
	}
	return nil
}

func ChangeStatus(username, passwd, data, goodsid, newaddr string, status uint16) error {
	reader := strings.NewReader(keystr)
	auth, err := bind.NewTransactor(reader, "123")
	if err != nil {
		fmt.Println("failed to NewTransactor ", err)
		return err
	}
	_, err = instance.ChangeGoodsStatus(auth, username, passwd, status, data, goodsid, newaddr)
	if err != nil {
		fmt.Println("failed to ChangeGoodsStatus ", err)
		return err
	}
	return nil
}

type TraceData struct {
	Addr      string
	Status    string
	Timestamp *big.Int
	Data      string
}

var DataStatus map[uint16]string = map[uint16]string{
	0: "产地库存中",
	1: "运输中",
	2: "商家库存中",
	3: "已售出",
}

func GetGoodsStatus(goodsid string) (string, error) {
	status, err := instance.GetStatus(nil, goodsid)
	if err != nil {
		fmt.Println("failed to GetStatus ", err)
		return "", err
	}

	return DataStatus[status], err
}

func GetGoodsMsg(goodsid string) ([]TraceData, error) {
	var traces []TraceData
	goodsTrace, err := instance.GetGoods(nil, goodsid)
	if err != nil {
		fmt.Println("failed to GetGoods ", err)
		return nil, err
	}
	var trace TraceData
	for _, v := range goodsTrace {
		trace.Addr = v.Addr
		trace.Data = v.Data
		trace.Timestamp = v.Timestamp
		trace.Status = DataStatus[v.Status]

		traces = append(traces, trace)
	}

	return traces, nil
}
