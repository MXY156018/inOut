package config

type Local struct {
	// 本地文件路径
	Path string `mapstructure:"path" json:"path" yaml:"path"`
	//本地域名
	Origin string `mapstructure:"origin" json:"origin" yaml:"origin"`
}
