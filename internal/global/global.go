package global

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/XxThunderBlast/thunder-api/internal/env"
)

var (
	Timer time.Time

	Env env.Env

	DB *mongo.Database
)
