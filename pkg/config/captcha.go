package config

type Captcha struct {
	// 验证码长度
	KeyLong int `mapstructure:"key-long" json:"keyLong" yaml:"key-long"`
	// 验证码宽度
	ImgWidth int `mapstructure:"img-width" json:"imgWidth" yaml:"img-width"`
	// 验证码高度
	ImgHeight int `mapstructure:"img-height" json:"imgHeight" yaml:"img-height"`
	//倾斜 0.7
	MaxSkew float64 `json:"maxSkew"`
	// 点的数量
	DotCount int `json:"dotCount"`
}
