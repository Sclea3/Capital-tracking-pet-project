package models

type Config struct {
	//Capital.com
	CAPITAL_API_KEY string `mapstructure:"CAPITAL_API_KEY"`
	CAPITAL_API_URL string `mapstructure:"CAPITAL_API_URL"`

	//Database
	DB_TYPE              string `mapstructure:"DB_TYPE"`
	DB_CONNECTION_STRING string `mapstructure:"DB_CONNECTION_STRING"`

	//Telegram
	TELEGRAM_BOT_TOKEN string `mapstructure:"TELEGRAM_BOT_TOKEN"`
	TELEGRAM_CHAT_ID   string `mapstructure:"TELEGRAM_CHAT_ID"`

	//App Settings
	POLL_INTERVAL_SECONDS uint8 `mapstructure:"POLL_INTERVAL_SECONDS"`
}
