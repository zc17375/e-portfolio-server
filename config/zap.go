package config

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

// Zap 日誌配置
type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                            // 日誌級別
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 日誌前綴
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                         // 日誌格式
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                  // 日誌文件夾
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       // 編碼級別
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // 堆棧名稱

	MaxAge       int  `mapstructure:"max-age" json:"max-age" yaml:"max-age"`                      // 日誌保留時間
	ShowLine     bool `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // 顯示行號
	LogInConsole bool `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // 是否輸出到控制台
}

func (z *Zap) ZapEncodeLevel() zapcore.LevelEncoder {
	switch {
	case z.EncodeLevel == "LowercaseLevelEncoder": // 小寫編碼器（默認）
		return zapcore.LowercaseLevelEncoder
	case z.EncodeLevel == "LowercaseColorLevelEncoder": // 小寫編碼器帶顏色
		return zapcore.LowercaseColorLevelEncoder
	case z.EncodeLevel == "CapitalLevelEncoder": // 大寫編碼器
		return zapcore.CapitalLevelEncoder
	case z.EncodeLevel == "CapitalColorLevelEncoder": // 大寫編碼器帶顏色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}


func (z *Zap) TransportLevel() zapcore.Level {
	z.Level = strings.ToLower(z.Level)
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
