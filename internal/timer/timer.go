package timer

import (
	"time"

	"github.com/XxThunderBlast/thunder-api/internal/global"
)

func InitTimer() {
	global.Timer = time.Now()
}

func TimeElapsedHR() float64 {
	return time.Since(global.Timer).Hours()
}

func TimeElapsedMin() float64 {
	return time.Since(global.Timer).Minutes()
}

func TimeElapsedSec() float64 {
	return time.Since(global.Timer).Seconds()
}
