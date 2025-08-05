package models

type KeyPair struct {
	Model
	Key   string `json:"key"`
	Value string `json:"value"`
}
