package global

import (
	"github.com/XxThunderBlast/thunder-api/internal/env"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var (
	Timer time.Time

	BaseKVPath string

	Env env.Env

	Db *mongo.Database
)
