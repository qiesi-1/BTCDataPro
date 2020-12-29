package ways

import (
	"BcRPCCode/constants"
	"BcRPCCode/entity"
	"BcRPCCode/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Method(method string, params... entity.Params) {
	rpcReq := entity.RPCRequest{}
	rpcReq.Id = time.Now().Unix()
	rpcReq.Jsonrpc = "2.0"
	rpcReq.Method = constants.GETNEWADDRESS
	//if params != nil {
	//	rpcReq.Params = params
	//}else {
	//
	//}

	//对结构体类型进行序列化
	rpcBytes, err := json.Marshal(&rpcReq)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Println("准备好的json格式的数据：", string(rpcBytes))

	//2,发送一个post类型的请求
	//客户端
	client := http.Client{}
	//实例化一个请求
	request, err := http.NewRequest("POST", constants.RPCURL, bytes.NewBuffer(rpcBytes))
	request.Header.Set("Encoding", "UTF-8")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Basic " + utils.Base64Str(constants.RPCUSER + ":" + constants.RPCPASSWORD))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//使用客户端发送请求
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//通过response获取响应数据
	code := response.StatusCode
	if code == 200 {
		resultBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		//fmt.Println("获取区块数量：")
		//json的反序列化
		var result entity.RPCResult
		err = json.Unmarshal(resultBytes, &result)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		//反序列化正常，没有出现错误
		fmt.Println(result.Result)
	} else {
		fmt.Println("抱歉，请求失败")
	}
}

