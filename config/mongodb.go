package config

type Mongo struct {
	Dbname   string `mapstructure:"db-name" json:"db-name" yaml:"db-name"` // 資料庫名稱
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Username string `mapstructure:"username" json:"username" yaml:"username"` // 資料庫使用者
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 資料庫密碼
}

func (m *Mongo) GetURI() string {
	return "mongodb://" + m.Username + ":" + m.Password + "@" + m.Host + ":" + m.Port
}
