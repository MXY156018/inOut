package config

// JWT 配置
type Jwt struct {
	// JWT 密钥
	Secret string `json:"secret"`
	// 超时时间
	ExpiresTime int64 `json:"expiresTime"`
	// 缓存时间
	BufferTime int64 `json:"bufferTime"`
}
