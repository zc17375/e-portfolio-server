package initialize

import (
	"os"

	"github.com/zc17375/e-portfolio-server/global"
	"github.com/zc17375/e-portfolio-server/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.EP_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.EP_DB
	err := db.AutoMigrate(
		// 新增相關Table Struct
		// system.SysApi{},
		model.User{},
		// system.SysBaseMenu{},
		model.JwtBlacklist{},
		// system.SysAuthority{},
		// system.SysDictionary{},
		// system.SysOperationRecord{},
		// system.SysAutoCodeHistory{},
		// system.SysDictionaryDetail{},
		// system.SysBaseMenuParameter{},
		// system.SysBaseMenuBtn{},
		// system.SysAuthorityBtn{},
		// system.SysAutoCode{},
		// system.SysExportTemplate{},
		// system.Condition{},

		// example.ExaFile{},
		// example.ExaCustomer{},
		// example.ExaFileChunk{},
		// example.ExaFileUploadAndDownload{},
	)
	if err != nil {
		global.EP_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.EP_LOG.Info("register table success")
}
