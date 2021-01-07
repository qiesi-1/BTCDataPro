package tools

import (
	"BTCDataPro/entity"
	"github.com/mapstructure-master"
)


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


//获取最新address
func GetNewAddress() (interface{}, error) {
	result, err := GetMsgByCommand("getnewaddress")
	if err != nil {
		return "", err
	}
	return result.Result, err
}