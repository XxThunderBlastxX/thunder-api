package service

import (
	"github.com/XxThunderBlast/thunder-api/domain"
)

type kvService struct {
	KVRepo domain.KVRepository
}

func NewKVService(kvRepo domain.KVRepository) domain.KVService {
	return &kvService{
		KVRepo: kvRepo,
	}
}

func (kv *kvService) GetValue(key string) (string, error) {
	return kv.KVRepo.GetValue(key)
}

func (kv *kvService) SetKeyValue(key string, value string) error {
	return kv.KVRepo.SetKeyValue(key, value)
}

func (kv *kvService) DeleteKey(key string) error {
	return kv.KVRepo.DeleteKey(key)
}

func (kv *kvService) ListKeys() ([]string, error) {
	return kv.KVRepo.ListKeys()
}
