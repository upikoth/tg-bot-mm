package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	BotToken           string `envconfig:"BOT_TOKEN" required:"true"`
	AuthFileDirName    string `envconfig:"YDB_AUTH_FILE_DIR_NAME" required:"true"`
	AuthFileName       string `envconfig:"YDB_AUTH_FILE_NAME" required:"true"`
	Dsn                string `envconfig:"YDB_DSN" required:"true"`
	NotificationChatID string `envconfig:"NOTIFICATION_CHAT_ID" required:"true"`
}

func New() (*Config, error) {
	config := &Config{}

	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}
