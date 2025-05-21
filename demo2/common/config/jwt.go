package config

// JWTConfig stores JWT configurations
type JWTConfig struct {
	Secret     string `mapstructure:"secret"`
	ExpireTime int    `mapstructure:"expire_time"` // in hours
}
