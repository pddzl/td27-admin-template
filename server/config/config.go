package config

type Server struct {
	JWT        JWT        `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap        Zap        `mapstructure:"zap" json:"zap" yaml:"zap"`
	RotateLogs RotateLogs `mapstructure:"rotateLogs" json:"rotateLogs" yaml:"rotateLogs"`
	System     System     `mapstructure:"system" json:"system" yaml:"system"`
	Mysql      Mysql      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis      Redis      `mapstructure:"redis" json:"redis" yaml:"redis"`
	Captcha    Captcha    `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
}
