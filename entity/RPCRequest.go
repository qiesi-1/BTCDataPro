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

type Params interface {
	ParamsInt() int64
	ParamsStr() string
}

type Param interface {
	Height() int
}

type GetBlockHeight struct{
    height int
}

func NewGetBlockByHeight(num int) []int {
	arr := []int{num}
	return arr
}

func (h GetBlockHeight) ParamsInt() int64 {
	return h.ParamsInt()
}
