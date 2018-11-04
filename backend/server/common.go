package server

import (
	"github.com/AleksandrKuts/youtubemeter-service/backend/config"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func init() {
	log = config.Logger
}
