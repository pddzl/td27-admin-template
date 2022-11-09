package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"server/config"
)

var (
	TD27_VP     *viper.Viper
	TD27_CONFIG config.Server
	TD27_LOG    *zap.Logger
	TD27_DB     *gorm.DB
)
