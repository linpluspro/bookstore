package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"bookstore/config"
)

// 数据库连接 配置管理器 日志记录器 配置信息
var (
	GVA_DB     *gorm.DB
	GVA_VP     *viper.Viper
	GVA_LOG    *zap.Logger
	GVA_CONFIG config.Server
)
