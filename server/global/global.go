package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"server/config"
)

var (
	TD27_VP     *viper.Viper
	TD27_CONFIG config.Server
	TD27_LOG    *zap.Logger
)
