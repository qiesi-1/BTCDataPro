package entity

type POOlinffo struct {
	Loaded bool `json:"loaded"`
	Size int64 `json:"size"`
	Bytes int64 `json:"bytes"`
	Usage int64 `json:"usage"`
	Maxmenpool int64 `json:"maxmenpool"`
	Mempoolminfee float64 `json:"mempoolminfee"`
	Minrelaytxfee float64 `json:"minrelaytxfee"`
}
