package config

type RedisConfig struct {
	DSN      string
	Password string
	DB       int
}

func loadRedisConfig() *RedisConfig {
	return &RedisConfig{
		DSN:      getString("REDIS_DSN", "localhost:6379"),
		Password: getString("REDIS_PASSWORD", ""),
		DB:       getInt("REDIS_DB", 0),
	}
}
