package kv

type Repository interface {
	Get(string) (string, error)
	Set(KeyValue) error
	Delete(string) error
}
