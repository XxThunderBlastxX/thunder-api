package global

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/XxThunderBlast/thunder-api/internal/gen/appconfig"
)

var (
	Timer time.Time

	DB *mongo.Database

	Config *appconfig.AppConfig
)
