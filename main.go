package main

import (
	"BcRPCCode/constants"
	"BcRPCCode/ways"
	"fmt"
)

func main() {
	//1,测试：获取区块数量
	fmt.Println("获取区块数量：")
	ways.Method(constants.GETBLOCKCOUNT)

	//2,测试：获取最新区块hash
	fmt.Println("最新区块的hash：")
	ways.Method(constants.GETBESTBLOCKHASH)

	//3，测试：生成一个新的比特币地址
	fmt.Println("新的比特币地址：")
	ways.Method(constants.GETNEWADDRESS)

	//4,测试：返回当前区块的难度值
	fmt.Println("当前区块的难度值：")
	ways.Method(constants.GETDIFFICULTY)

	//5，测试：返回当前节点的连接数
	fmt.Println("当前节点连接数：")
	ways.Method(constants.GETCONNECTIONCOUNT)

	//6,测试：获取某个高度的hash
	fmt.Println("特定高度区块的hash：")
	ways.Method(constants.GETBLOCKBYHEIGHT)

	//7，测试：获取内存池
	fmt.Println("内存池为：")
	ways.Method(constants.GETMEMPOOLINFO)
}
