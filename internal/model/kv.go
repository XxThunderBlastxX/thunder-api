package model

type KVRequest struct {
	Base64        bool   `json:"base64,omitempty"`
	Expiration    int    `json:"expiration,omitempty"`
	ExpirationTTL int    `json:"expiration_ttl,omitempty"`
	Key           string `json:"key"`
	Value         string `json:"value"`
}

type KVResponse struct {
	Error      string               `json:"error"`
	Messages   []string             `json:"messages"`
	Result     []kvResponseResult   `json:"result"`
	Success    bool                 `json:"success"`
	ResultInfo kvResponseResultInfo `json:"result_info"`
}

type kvResponseResult struct {
	Expiration int               `json:"expiration"`
	MetaData   map[string]string `json:"metadata"`
	Message    string            `json:"message"`
}

type kvResponseResultInfo struct {
	Count  int    `json:"count"`
	Cursor string `json:"cursor"`
}
