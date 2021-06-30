package config

type Redis struct {
	Db          int    `mapstructure:"db" json:"db" yaml:"db"`
	Addr        string `mapstructure:"addr" json:"addr" yaml:"addr"`
	RequirePass bool   `mapstructure:"require-pass" json:"requirePass" yaml:"require-pass"`
	Password    string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdle     int    `mapstructure:"max-idle" json:"maxIdle" yaml:"max-idle"`
	MaxActive   int    `mapstructure:"max-active" json:"maxActive" yaml:"max-active"`
	IdleTimeout int    `mapstructure:"idle-timeout" json:"idleTimeout" yaml:"idle-timeout"`
	Wait        bool   `mapstructure:"wait" json:"wait" yaml:"wait"`
}
