package config

type Server struct {
	Addr           int `yaml:"addr"`
	ReadTimeout    int `mapstructure:"read-timeout" yaml:"read-timeout"`
	WriteTimeout   int `mapstructure:"write-timeout" yaml:"write-timeout"`
	MaxHeaderBytes int `mapstructure:"max-header-bytes" yaml:"max-header-bytes"`
}
