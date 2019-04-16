package dorker

import (
	"github.com/gocaio/goca/dorker/logger"
)

// LogMeIn logs the scrapped files
func LogMeIn() {
	logger.NewGame().Start()
}
