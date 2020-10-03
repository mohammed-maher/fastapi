package config

type SMSConfig struct {
	Sender string
	Token  string
	Expiry string
}

func LoadSMSConfig() *SMSConfig {
	return &SMSConfig{
		Sender: getString("SMS_SENDER", ""),
		Token:  getString("SMS_TOKEN", ""),
		Expiry: getString("SMS_EXPIRY", "900"),
	}
}
