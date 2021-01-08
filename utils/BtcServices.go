package main

import (
	"BtcRPC/entity"
	"BtcRPC/mapstructure"
	"fmt"
)

func main() {
	Getblockinfo()
	GetBlocks()
}

func Getblockinfo(){
	getblockchaininfo,err :=GetBlockChainInfo()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("最新区块链信息为：",getblockchaininfo.Chain)
	fmt.Println("区块链的难度为：",getblockchaininfo.Difficulty)
	fmt.Println("区块链上工作量总值为：",getblockchaininfo.Chainwork)
	fmt.Println("区块链的头部：",getblockchaininfo.Headers)
	fmt.Println("区块链的软分叉：",getblockchaininfo.Softforks)
}

func GetBlocks()  {
	const address  = "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f"
	getblock,err := GetBlockByHash(address)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("dasd",getblock.Tx)
}
//
//比特币节点命令 getblockchaininfo 的封装函数
func GetBlockChainInfo() (*entity.BlockChainInfo, error) {
	result, err := GetAllCommand("getblockchaininfo")
	if err != nil {
		return nil, err
	}

	var blockchain entity.BlockChainInfo
	err = mapstructure.Decode(result.Result, &blockchain)
	if err != nil{
		return nil, err
	}
	return &blockchain,err

}

//比特币节点命令 getblockcount 的封装函数
func GetBlockCount() (interface{}, error) {
	result, err := GetAllCommand("getblockcount")
	if err != nil {
		return "", err
	}
	return result.Result, err
}

//比特币节点命令 getdifficult 的封装函数
func GetDifficult()(interface{}, error)  {
	result, err := GetAllCommand("getdifficult")
	if err != nil {
		return nil,err
	}
	return result.Result,nil
}

//比特币节点命令 getbestblockhash 的封装函数
func GetBestBlockHash() (interface{}, error) {
	result, err := GetAllCommand("getbestblockhash")
	if err != nil {
		return "", err
	}
	return result.Result, err
}

//比特币节点命令 getblockhash 的封装函数
func GetBlockHashByHeight(height int64) (interface{}, error) {
	result, err := GetAllCommand("getblockhash",height)
	if err != nil {
		return "", err
	}
	return result.Result, err
}

//比特币节点命令 getblock 的封装函数
func GetBlockByHash(hash string) (*entity.GetBlock,error) {
	result, err:= GetAllCommand("getblock", hash)
	if err != nil {
		return nil, err
	}
	var block entity.GetBlock
	err = mapstructure.Decode(result.Result, &block)
	if err != nil {
		return nil,err
	}
	return &block, nil
}

//比特币节点命令 getblockheader hash的封装函数
func GetBlockHeaderByHash(hash string) (*entity.RPCResult,error) {
	result, err:= GetAllCommand("getblockheader", hash)
	if err != nil {
		return nil, err
	}
	var block entity.RPCResult
	err = mapstructure.Decode(result.Result, &block)
	if err != nil {
		return nil,err
	}
	return &block, nil
}
