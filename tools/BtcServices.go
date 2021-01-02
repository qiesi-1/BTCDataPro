package tools

import (
	"BTCDataPro/entity"
	"BTCDataPro/models"
	"fmt"
	"github.com/mapstructure-master"
)

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
	info := models.Info{
		Hash:      "",
		Height:    0,
		Difficult: 0,
		Size:      0,
		Count:     0,
		Chain:     "",
	}
	//保存到数据库中
	_, err=info.AddRpcInfo()
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
	result, err := GetMsgByCommand("getblockchaininfo")
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
	result, err := GetMsgByCommand("getblockcount")
	if err != nil {
		return "", err
	}
	//存入数据

	//db.O.Insert(&result.Result)
	return result.Result, err
}
//比特币节点命令 getdifficult 的封装函数
func GetDifficult()(interface{}, error)  {
	result, err := GetMsgByCommand("getdifficult")
	if err != nil {
		return nil,err
	}
	return result.Result,nil
}
//比特币节点命令 getbestblockhash 的封装函数
func GetBestBlockHash() (interface{}, error) {
	result, err := GetMsgByCommand("getbestblockhash")
	if err != nil {
		return "", err
	}
	return result.Result, err
}

//比特币节点命令 getblockhash 的封装函数
func GetBlockHashByHeight(height int64) (interface{}, error) {
	result, err := GetMsgByCommand("getblockhash",height)
	if err != nil {
		return "", err
	}
	return result.Result, err
}

//比特币节点命令 getblock 的封装函数
func GetBlockByHash(hash string) (*entity.GetBlock,error) {
	result, err:= GetMsgByCommand("getblock", hash)
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
	result, err:= GetMsgByCommand("getblockheader", hash)
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
