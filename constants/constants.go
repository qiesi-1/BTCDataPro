package constants

//常量
const RPCURL = "http://127.0.0.1:8332"
const RPCUSER = "xiaoyang"
const RPCPASSWORD = "714736690"

//地址类型
const ADDRESS_LEGACY = "legacy"
const ADDRESS_P2SH_SEGWIT = "p2sh-segwit"
const ADDRESS_BACH32 = "bach32"

//命令
/**
获取区块数量
 */
const GETBLOCKCOUNT = "getblockcount"
/**
获取最新区块的hash
 */
const GETBESTBLOCKHASH = "getbestblockhash"
/**
获取一个新的地址
 */
const GETNEWADDRESS = "getnewaddress"
/**
获取当前区块难度
 */
const GETDIFFICULTY = "getdifficulty"
/**
返回当前节点的连接数
 */
const GETCONNECTIONCOUNT = "getconnectioncount"
/**
获取某个高度的hash
 */
const GETBLOCKBYHEIGHT = "getblockhash"
/**
通过区块的哈希值获取区块数据
 */
const GETBLOCKBYHASH = "getblock"
/**
获取内存池数据
 */
const GETMEMPOOLINFO = "getmempoolinfo"
