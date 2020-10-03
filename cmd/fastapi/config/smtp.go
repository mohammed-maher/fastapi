package config

type SmtpConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

func LoadSmtpConfig() *SmtpConfig {
	return &SmtpConfig{
		Host:     getString("SMTP_HOST", "mail.example.com"),
		Port:     getInt("SMTP_PORT", 587),
		Username: getString("SMTP_USERNAME", "info@example.com"),
		Password: getString("SMTP_PASSWORD", ""),
	}
}
