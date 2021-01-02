package entity

type RPCRequest struct {
	Id      int64         `json:"id"`
	Method  string        `json:"method"`
	Jsonrpc string        `json:"jsonrpc"`
	Params  []interface{} `json:"params"`
}

type RPCResult struct {
	Id     int64       `json:"id"`
	Error  interface{} `json:"error"`
	Result interface{} `json:"result"`
}

type GetBlock struct {
	Hash              interface{} `json:"hash"`
	Confirmations     int64       `json:"confirmations"`
	Size              int64       `json:"size"`
	Strippedsize      int64       `json:"strippedsize"`
	Weight            int64       `json:"weight"`
	Height            int64       `json:"height"`
	Version           int         `json:"version"`
	VersionHex        string      `json:"version_hex"`
	Tx                []string    `json:"tx"`
	Time              int64       `json:"time"`
	Mediantime        int64       `json:"mediantime"`
	Nonce             int64       `json:"nonce"`
	Bits              string      `json:"bits"`
	Difficulty        int64       `json:"difficulty"`
	Chainwork         string      `json:"chainwork"`
	Ntx               int64       `json:"ntx"`
	NexblockHash      string      `json:"nexblock_hash"`
	Previousblockhash string      `json:"previousblockhash"`
	Nextblockhash     string      `json:"nextblockhash"`
	Result            interface{} `json:"result"`
}

type BlockChainInfo struct {
	Chain                string `json:"chain"`
	Blocks               int64  `json:"blocks"`
	Headers              int64  `json:"headers"`
	Bestblockhash        string `json:"bestblockhash"`
	Difficulty           int64  `json:"difficulty"`
	Mediantime           int64  `json:"mediantime"`
	Verificationprogress int64  `json:"verificationprogress"`
	Initialblockdownload bool   `json:"initialblockdownload"`
	Chainwork            string `json:"chainwork"`
	Size_on_disk         int64  `json:"size_on_disk"`
	Pruned               bool  `json:"pruned"`
	Pruneheight          int64  `json:"pruneheight"`
	Automatic_pruning    bool   `json:"automatic_pruning"`
	Prune_target_size    int64  `json:"prune_target_size"`
	Softforks            Bip     `json:"softforks"`
	Warnings             string `json:"warnings"`
}

type Bip struct {
	Bip34  a `json:"bip_34"`
	Bip66  a `json:"bip_66"`
	Bip65  a `json:"bip_65"`
	Csv    a `json:"csv"`
	Segwit a `json:"segwit"`
}

type a struct {
	Types string `json:"type"`
	Active bool  `json:"active"`
	Height int64 `json:"height"`
}

