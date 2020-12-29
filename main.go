package main

import (
	"BcRPCCode/constants"
	"BcRPCCode/ways"
)

func main() {
	//1,测试：获取区块数量
	ways.Method(constants.GETBLOCKCOUNT)

	//2,测试：获取最新区块hash
	ways.Method(constants.GETBESTBLOCKHASH)

	//3，测试：生成一个新的比特币地址
	ways.Method(constants.GETNEWADDRESS)

	//4,测试：返回当前区块的难度值
	ways.Method(constants.GETDIFFICULTY)

	//5，测试：返回当前节点的连接数
	ways.Method(constants.GETCONNECTIONCOUNT)

	//6,测试：获取某个高度的hash
	ways.Method(constants.GETBLOCKBYHEIGHT)

	//7，测试：获取内存池
	ways.Method(constants.GETMEMPOOLINFO)

	//8,测试:通过区块哈希获取区块数据
    ways.Method(constants.GETBESTBLOCKHASH)
}
