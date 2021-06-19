package config

type JWT struct {
	Secret  string `mapstructure:"secret" json:"secret" yaml:"secret"`
	Ttl int64  `mapstructure:"ttl" json:"ttl" yaml:"ttl"`
}
