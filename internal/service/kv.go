package service

import (
	"context"
	"fmt"
	"io"

	"github.com/XxThunderBlastxX/thunder-api/internal/config"
	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/kv"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

type KVService struct {
	cfKVService *kv.NamespaceService
	namespaceID string
	accountID   string
}

func NewKVService(cfg *config.CloudflareConfig) *KVService {
	// Initialize Cloudflare client with API token
	client := cloudflare.NewClient(
		option.WithAPIToken(cfg.Token),
	)

	cf := client.KV.Namespaces
	namespaceID := cfg.KVNamespaceID
	accntID := cfg.AccountID

	return &KVService{
		cfKVService: cf,
		namespaceID: namespaceID,
		accountID:   accntID,
	}
}

func (s *KVService) GetValue(key string) (string, error) {
	value, err := s.cfKVService.Values.Get(
		context.TODO(),
		s.namespaceID,
		key,
		kv.NamespaceValueGetParams{
			AccountID: cloudflare.F(s.accountID),
		},
	)
	if err != nil {
		return "", err
	}
	defer value.Body.Close()

	res, err := io.ReadAll(value.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(res), nil
}

func (s *KVService) SetKeyValue(key string, value string) error {
	_, err := s.cfKVService.Values.Update(
		context.TODO(),
		s.namespaceID,
		key,
		kv.NamespaceValueUpdateParams{
			AccountID: cloudflare.F(s.accountID),
			Value:     cloudflare.F(value),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to set key-value pair: %w", err)
	}

	return nil
}

func (s *KVService) DeleteKey(key string) error {
	_, err := s.cfKVService.Values.Delete(
		context.TODO(),
		s.namespaceID,
		key,
		kv.NamespaceValueDeleteParams{
			AccountID: cloudflare.F(s.accountID),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to delete key: %w", err)
	}

	return nil
}

func (s *KVService) ListKeys() ([]string, error) {
	page, err := s.cfKVService.Keys.List(
		context.TODO(),
		s.namespaceID,
		kv.NamespaceKeyListParams{
			AccountID: cloudflare.F(s.accountID),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list keys: %w", err)
	}

	var keys []string
	for _, key := range page.Result {
		keys = append(keys, key.Name)
	}

	return keys, nil
}
