package main

import (
	"BtcRPC/entity"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)
//获取比特币节点中的 rpcUrl 、rpcAuthorization
//var rpcUrl= beego.AppConfig.String("RPCURL")
//var rpcAuthorization = beego.AppConfig.String("RPCAuthorization")
var rpcUrl = "http://127.0.0.1:8086"
var rpcAuthorization = "ztc:birthday"
//获取比特币节点的字符串形式
/*
method： 调用的具体命令
parms：参数
return 将rpc请求所需的数据打包并序列化为json格式
 */
func GetBTCJsonDeal(method string, parms ...interface{}) string {
	rpcReq := entity.RPCRequest{}
	rpcReq.Jsonrpc = "2.0"
	rpcReq.Id = time.Now().Unix()
	rpcReq.Method = method
	if parms != nil {
		rpcReq.Params = parms
	}
	//序列化
	objStr, err := json.Marshal(rpcReq)
	if err != nil {
		return ""
	}
	return string(objStr)
}
//将 json格式的数据 通过rpc请求方式向比特币节点获取节点数据
/*
jsonDeal： json格式的请求数据
BTCResult： rpc请求比特币节点的结果集
error :请求数据或解析遇到的错误
 */
func Excute(jsonStr string) (*entity.RPCResult, error) {
	clinet := &http.Client{}

	req, err := http.NewRequest("POST", rpcUrl, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Encoding", "UTF-8")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(rpcAuthorization)))

	response, err := clinet.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	//反序列化
	var rpcResult entity.RPCResult
	err = json.Unmarshal(body, &rpcResult)
	if err != nil {
		return nil, err
	}

	return &rpcResult, nil
}
/*btc命令调用封装函数 命令 [参数1，参数2 ...]
	method： 比特币节点具体命令
	parms ：命令对应的具体参数
	return：比特币 Result
 */
func GetAllCommand(method string, parms ...interface{}) (*entity.RPCResult, error) {
	jsonStr := GetBTCJsonDeal(method, parms...)
	return Excute(jsonStr)
}

