package types

type kvResponseResult struct {
	Expiration int               `json:"expiration"`
	MetaData   map[string]string `json:"metadata"`
	Message    string            `json:"message"`
}

type kvResponseResultInfo struct {
	Count  int    `json:"count"`
	Cursor string `json:"cursor"`
}

type KVResponse struct {
	Error      string               `json:"error"`
	Messages   []string             `json:"messages"`
	Result     []kvResponseResult   `json:"result"`
	Success    bool                 `json:"success"`
	ResultInfo kvResponseResultInfo `json:"result_info"`
}
