package entity

type RPCRequest struct {
	/**
	“id”，编号
	“meyhod”，功能方法或命令
	“jsonrpc”：rpc版本2.0
	“params”，【携带的参数，数组形式】
	*/
	Id      int64       `json:"id"`
	Method  string      `json:"method"`
	Jsonrpc string      `json:"jsonrpc"`
	Params  interface{} `json:"params"`
}

//参数标准
type Params interface {
	ParamsInt() int64
	ParamsStr() string
}

//获取某个特定高度的区块数据
type GetBlockByHeight struct{}

func NewGetBlockByHeight(num int) []int {
	arr := []int{num}
	return arr
}

func (h GetBlockByHeight) ParamsInt() int64 {
	return h.ParamsInt()
}

//根据哈希值获取区块数据
type Getblock struct {}

func NewGetblock(s string) []string {
	arr := []string{s}
	return arr
}

func (s Getblock) ParamsStr() string {
	return s.ParamsStr()
}
