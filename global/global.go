package global

import (
	"fmt"
	"github.com/XxThunderBlast/thunder-api/internal/env"
	"time"
)

var (
	Timer time.Time

	BaseKVPath = fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/storage/kv/namespaces/%s", Env.CFAccountId, Env.KvNamespaceId)

	Env env.Env
)
