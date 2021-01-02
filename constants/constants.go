package constants

const RPCURL = "http://127.0.0.1:8332"
const RPCUSER = "user"
const RPCPASSWDRD = "pwd"

//地址类型
const ADDRESS_LEGACY = "legacy"

const ADDRESS_P2SH_SEGWIT = "p2sh-segwit"

const ADDRESS_BECH32 = "bach32"

//rpc服务命令
//获取区块总数
const GETBLOCKCOUNT = "getblockcount"
//获取最新区块的哈希
const GETBESTBLOCKHASH = "getbestblockhash"
//生成一个新的比特币地址
const GETNEWADDRESS = "getnewaddress"
//获取当前比特币网络中的区块难度
const GETDIFFICULTY = "getdifficulty"
//通过区块的哈希值去获取区块数据
const GETBLOCKBYHASH = "getblock"
//获取区块链信息
const GETBLOCKCHAININFO = "getblockchaininfo"
//获取钱包的信息
const GETWALLETINFO = "getwalletinfo"
//获取某个高度的哈希值
const GETBLOCKHASHBYHEIGHT = "getblockhash"
//获取所有命令
const GETBLOCKCOMMAND = "help"
//获取当前服务器运行的秒数
const GETBLOCKSERVERTIME = "uptime"
//设置交易比特币的费用每kb
const GETSETTXFEE = "settxfee"
//设置与地址关联的标签
const GETSETLABEL = "setlabel"
//停止由RPC调用触发的当前钱包重新扫描。
const GETABORTRESCAN = "abortrescan"
//验证已签名的消息
const GETVERIFYMESSAGE = "verifymessage"
//请求向所有其他节点发送一个ping，以测量ping时间
const GETPING = "ping"

