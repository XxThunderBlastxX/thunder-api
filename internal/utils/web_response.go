package utils

type WebResponse struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitzero,omitempty"`
}
