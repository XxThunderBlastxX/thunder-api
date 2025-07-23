package cloudflare

import (
	"context"

	"github.com/XxThunderBlastxX/thunder-api/internal/config"
	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/kv"
)

type CloudflareKV struct {
	*kv.NamespaceService
	client    *cloudflare.Client
	namespace string
	accountID string
}

func NewCloudflareKV(cfg config.CloudflareConfig) *CloudflareKV {
	client := NewClient(cfg)
	namespaceService := client.KV.Namespaces
	ns := cfg.KVNamespaceID
	aid := cfg.AccountID

	return &CloudflareKV{
		client:           client,
		NamespaceService: namespaceService,
		namespace:        ns,
		accountID:        aid,
	}
}

func (kc *CloudflareKV) GetBulkNamespaceKey(ctx context.Context, keys []string) (*kv.NamespaceBulkGetResponse, error) {
	res, err := kc.BulkGet(ctx, kc.namespace, kv.NamespaceBulkGetParams{
		AccountID: cloudflare.F(kc.accountID),
		Keys:      cloudflare.F(keys),
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (kc *CloudflareKV) WriteBulkNamespaceKV(ctx context.Context, body []kv.NamespaceBulkUpdateParamsBody) (*kv.NamespaceBulkUpdateResponse, error) {
	res, err := kc.BulkUpdate(ctx, kc.namespace, kv.NamespaceBulkUpdateParams{
		AccountID: cloudflare.F(kc.accountID),
		Body:      body,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (kc *CloudflareKV) DeleteBulkNamespaceKey(ctx context.Context, keys []string) (*kv.NamespaceBulkDeleteResponse, error) {
	res, err := kc.BulkDelete(ctx, kc.namespace, kv.NamespaceBulkDeleteParams{
		AccountID: cloudflare.F(kc.accountID),
		Body:      keys,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
