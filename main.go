package main

import (
	"BcRPCCode1/constants"
	"BcRPCCode1/ways"
	"fmt"
)



func main() {
	//1,测试：获取区块数量
	fmt.Print("获取区块数量：")
	ways.Method(constants.GETBLOCKCOUNT)

	//2,测试：获取最新区块hash
	fmt.Print("最新区块的hash：")
	ways.Method(constants.GETBESTBLOCKHASH)

	//3，测试：生成一个新的比特币地址
	fmt.Print("新的比特币地址：")
	ways.Method(constants.GETNEWADDRESS)

	//4,测试：返回当前区块的难度值
	fmt.Print("当前区块的难度值：")
	ways.Method(constants.GETDIFFICULTY)

	//5，测试：返回当前节点的连接数
	fmt.Print("当前节点连接数：")
	ways.Method(constants.GETCONNECTIONCOUNT)

	//,测试：获取某个高度的hash
	//u := user{}
	//sendNotification(&u) // 传递的是指针
	fmt.Println("特定高度区块的hash：")
	ways.Method(constants.GETBLOCKBYHEIGHT)

	fmt.Print("当前ping值：")
	ways.Method(constants.PING)

	fmt.Print("难度：")
	ways.Method(constants.SACEMEMPOOL)

}
