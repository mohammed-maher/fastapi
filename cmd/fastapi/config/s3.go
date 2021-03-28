package config

type S3Config struct {
	Endpoint     string
	AccessKey    string
	SecretKey    string
	UploadBucket string
	SSLMode      bool
}

func LoadS3Config() *S3Config {
	return &S3Config{
		Endpoint:  getString("S3_ENDPOINT", "127.0.0.1"),
		AccessKey: getString("S3_ACCESS_KEY", ""),
		SecretKey: getString("S3_SECRET_KEY", ""),
		UploadBucket: getString("S3_UPLOAD_BUCKET", ""),
		SSLMode:   getBool("S3_ENABLE_SSL", false),
	}
}
