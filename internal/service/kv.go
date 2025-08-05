package service

type KVService struct{}

func NewKVService() *KVService {
	return &KVService{}
}

func (kv *KVService) GetValue(key string) (string, error) {
	panic("GetValue method not implemented")
}

func (kv *KVService) SetKeyValue(key string, value string) error {
	panic("SetKeyValue method not implemented")
}

func (kv *KVService) DeleteKey(key string) error {
	panic("DeleteKey method not implemented")
}

func (kv *KVService) ListKeys() ([]string, error) {
	panic("ListKeys method not implemented")
}
