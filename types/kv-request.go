package types

type KVRequest struct {
	Base64        bool   `json:"base64,omitempty"`
	Expiration    int    `json:"expiration,omitempty"`
	ExpirationTTL int    `json:"expiration_ttl,omitempty"`
	Key           string `json:"key"`
	Value         string `json:"value"`
}
