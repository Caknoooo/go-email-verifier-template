package config

import "github.com/spf13/viper"

type Email struct {
	Host       string `mapstructure:"SMTP_HOST"`
	Port       int    `mapstructure:"SMTP_PORT"`
	SenderName string `mapstructure:"SMTP_SENDER_NAME"`
	Mail       string `mapstructure:"SMTP_MAIL"`
	Password   string `mapstructure:"SMTP_PASSWORD"`
}

func NewEmailConfig() (*Email, error) {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.AutomaticEnv()

	var email Email
	if err := viper.Unmarshal(&email); err != nil {
		return nil, err
	}

	return &email, nil
}
