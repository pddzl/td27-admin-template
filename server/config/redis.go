package config

type Redis struct {
	DB       string `mapstructure:"db" json:"db" yaml:"db"`
	Path     string `mapstructure:"path" json:"path" yaml:"path"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
