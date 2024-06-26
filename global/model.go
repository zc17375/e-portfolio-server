package global

import (
	"time"

	"gorm.io/gorm"
)

type EP_MODEL struct {
	ID        uint           `gorm:"primarykey"` // 主鍵ID
	CreatedAt time.Time      // 創建時間
	UpdatedAt time.Time      // 更新時間
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 刪除時間
}
