package domain

type CFTurnstileToken struct {
	Secret   string `json:"secret"`
	Response string `json:"response"`
	RemoteIp string `json:"remoteip"`
}
