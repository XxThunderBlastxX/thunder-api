package cloudflare

import (
	"github.com/XxThunderBlastxX/thunder-api/internal/config"
	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func NewClient(cfg config.CloudflareConfig) *cloudflare.Client {
	client := cloudflare.NewClient(
		option.WithAPIToken(cfg.Token),
	)

	return client
}
