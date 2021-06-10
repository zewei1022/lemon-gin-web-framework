package config

type Database struct {
	Type         string `mapstructure:"type" json:"type" yaml:"type"`
	Hostname     string `mapstructure:"hostname" json:"hostname" yaml:"hostname"`
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	DbName       string `mapstructure:"db-name" json:"dbName" yaml:"db-name"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
	LogZap       string `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`
}