package config

type Server struct {
	Addr           int `mapstructure:"addr" json:"addr" yaml:"addr"`
	ReadTimeout    int `mapstructure:"read-timeout" json:"readTimeout" yaml:"read-timeout"`
	WriteTimeout   int `mapstructure:"write-timeout" json:"writeTimeout" yaml:"write-timeout"`
	MaxHeaderBytes int `mapstructure:"max-header-bytes" json:"maxHeaderBytes" yaml:"max-header-bytes"`
}
