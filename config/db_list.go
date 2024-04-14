package config

type DsnProvider interface {
	Dsn() string
}

// GeneralDB 也被 Pgsql 和 Mysql 使用
type GeneralDB struct {
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Port         string `mapstructure:"port" json:"port" yaml:"port"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`       // 高級配置
	Dbname       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`    // 資料庫名稱
	Username     string `mapstructure:"username" json:"username" yaml:"username"` // 資料庫使用者
	Password     string `mapstructure:"password" json:"password" yaml:"password"` // 資料庫密碼
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Engine       string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`        //數據庫引擎，默認InnoDB
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // 是否開啟Gorm全局日誌
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 閒置時最大連接數
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打開到數據庫的最大連接數
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`                   // 是否開啟全局禁用是否開啟全局複數，true表示開啟
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // 是否通過zap寫入日誌文件
}
