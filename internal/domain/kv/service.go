package kv

import "github.com/XxThunderBlastxX/thunder-api/internal/infrastructure/cloudflare"

type Service interface {
	GetValue(string) (string, error)
	ListKeys() ([]string, error)
	SetKey(string, string) error
	DeleteKey(string) error
}

type service struct {
	cfKVService cloudflare.CloudflareKV
}

func NewService(cfKVService cloudflare.CloudflareKV) Service {
	return &service{
		cfKVService: cfKVService,
	}
}

func (s *service) GetValue(key string) (string, error) {
	// TODO: Implement the logic
	panic("Method not implemented")
}

func (s *service) ListKeys() ([]string, error) {
	// TODO: Implement the logic
	panic("Method not implemented")
}

func (s *service) SetKey(key, value string) error {
	// TODO: Implement the logic
	panic("Method not implemented")
}

func (s *service) DeleteKey(key string) error {
	// TODO: Implement the logic
	panic("Method not implemented")
}
