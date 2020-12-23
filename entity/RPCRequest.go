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

type GetBlockHeight struct{}

func NewGetBlockByHeight(i interface{}) interface{} {
	return i
}

func (h GetBlockHeight) ParamsInt() int64 {
	return h.ParamsInt()
}
