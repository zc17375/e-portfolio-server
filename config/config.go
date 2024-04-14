package config

type Server struct {
	Jwt JWT `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	// Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	// Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	System System `mapstructure:"system" json:"system" yaml:"system"`

	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`

	// mongo
	Mongo Mongo `mapstructure:"mongo" json:"mongo" yaml:"mongo"`


	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`

	// // added by mohamed hassan to support multi-language
	// Language Language `mapstructure:"language" json:"language" yaml:"language"`
}
