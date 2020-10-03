package config

type JWTConfig struct {
	AccessSecret  string
	RefreshSecret string
}

func LoadJWTConfig() *JWTConfig {
	return &JWTConfig{
		AccessSecret:  getString("JWT_ACCESS_SECRET", ""),
		RefreshSecret: getString("JWT_REFRESH_SECRET", ""),
	}
}
