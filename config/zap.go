package config

type Zap struct {
	Level       string `mapstructure:"level" json:"level" yaml:"level"`
	EncodeLevel string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`
	Director    string `mapstructure:"director" json:"director"  yaml:"director"`
	MaxSize     int    `mapstructure:"max-size" json:"maxSize"  yaml:"max-size"`
	MaxBackups  int    `mapstructure:"max-backups" json:"MaxBackups"  yaml:"max-backups"`
	MaxAge      int    `mapstructure:"max-age" json:"MaxAge"  yaml:"max-age"`
	Compress    bool   `mapstructure:"compress" json:"Compress"  yaml:"compress"`
}
