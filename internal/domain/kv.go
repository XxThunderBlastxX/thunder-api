package domain

type KVRepository interface {
	GetValue(key string) (string, error)
	SetKeyValue(key string, value string) error
	DeleteKey(key string) error
	ListKeys() ([]string, error)
}

type KVService interface {
	GetValue(key string) (string, error)
	SetKeyValue(key string, value string) error
	DeleteKey(key string) error
	ListKeys() ([]string, error)
}
