package global

import (
	"github.com/spf13/viper"
	"github.com/zc17375/e-portfolio-server/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	EP_DB     *gorm.DB
	EP_DBList map[string]*gorm.DB
	// GVA_REDIS  *redis.Client
	EP_CONFIG config.Server
	EP_VP     *viper.Viper
	EP_LOG                 *zap.Logger
	// GVA_Timer               timer.Timer = timer.NewTimerTask()
	// GVA_Concurrency_Control             = &singleflight.Group{}

	// BlackCache local_cache.Cache
	// lock       sync.RWMutex

	// GVA_TRANSLATOR translate.Translator // added by mohamed hassan to support multilanguage
)
