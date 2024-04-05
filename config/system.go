package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                // 環境值
	DbType        string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`    // 數據庫類型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType       string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"` // Oss类型
	RouterPrefix  string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	Port          int    `mapstructure:"port" json:"port" yaml:"port"` // 端口值
	LimitCountIP  int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP   int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
}
