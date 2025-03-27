package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"CapitalParser/internal/models"

	"github.com/joho/godotenv"
)

// LoadConfig loads configuration from .env file
func LoadConfig() (models.Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	capitalApiKey := os.Getenv("CAPITAL_API_KEY")
	capitalUsername := os.Getenv("CAPITAL_USERNAME")
	capitalPassword := os.Getenv("CAPITAL_PASSWORD")

	dbType := os.Getenv("DB_TYPE")
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")

	telegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	telegramChatId := os.Getenv("TELEGRAM_CHAT_ID")

	pollIntervalSeconds, err := strconv.Atoi(os.Getenv("POLL_INTERVAL_SECONDS"))

	if pollIntervalSeconds == 0 || err != nil {
		pollIntervalSeconds = 10
		log.Println("POLL_INTERVAL_SECONDS is not set in .env file or set as 0. Using default value 10")
	}

	//Check if all environment variables are set
	err = checkEnvVariables(capitalApiKey, capitalUsername, capitalPassword, dbType, dbConnectionString, telegramBotToken, telegramChatId, pollIntervalSeconds)
	if err != nil {
		return models.Config{}, err
	}

	return models.Config{
		CAPITAL_API_KEY:       capitalApiKey,
		CAPITAL_USERNAME:      capitalUsername,
		CAPITAL_PASSWORD:      capitalPassword,
		DB_TYPE:               dbType,
		DB_CONNECTION_STRING:  dbConnectionString,
		TELEGRAM_BOT_TOKEN:    telegramBotToken,
		TELEGRAM_CHAT_ID:      telegramChatId,
		POLL_INTERVAL_SECONDS: uint8(pollIntervalSeconds),
	}, nil
}

func checkEnvVariables(apiKey, username, password, dbType, dbConString, botToken, chatId string, pollInterval int) error {
	envVars := map[string]string{
		"CAPITAL_API_KEY":       apiKey,
		"CAPITAL_USERNAME":      username,
		"CAPITAL_PASSWORD":      password,
		"DB_TYPE":               dbType,
		"DB_CONNECTION_STRING":  dbConString,
		"TELEGRAM_BOT_TOKEN":    botToken,
		"TELEGRAM_CHAT_ID":      chatId,
		"POLL_INTERVAL_SECONDS": strconv.Itoa(pollInterval),
	}
	for key, value := range envVars {
		if value == "" {
			return fmt.Errorf("environment variable %s is not set", key)
		}
	}
	return nil
}
