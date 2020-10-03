package config

type S3Config struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	SSLMode   bool
}

func LoadS3Config() *S3Config {
	return &S3Config{
		Endpoint:  getString("S3_ENDPOINT", "127.0.0.1"),
		AccessKey: getString("S3_ACCESS_KEY", ""),
		SecretKey: getString("S3_SECRET_KEY", ""),
		SSLMode:   getBool("S3_ENABLE_SSL", false),
	}
}
