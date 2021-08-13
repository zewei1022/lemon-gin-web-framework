package config

type Mongodb struct {
	Uri string `mapstructure:"uri" json:"uri" yaml:"uri"`
	ConnectTimeout int `mapstructure:"connect-timeout" json:"connectTimeout" yaml:"connect-timeout"`
	PingTimeOut int `mapstructure:"ping-timeout" json:"pingTimeout" yaml:"ping-timeout"`
}
