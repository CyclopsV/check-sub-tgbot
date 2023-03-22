package configs

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	TgBotToken string `json:"TG_BOT_TOKEN"`
	ChannelID  int64    `json:"CHANNEL_ID"`
	ChatID     int64    `json:"CHAT_ID"`
}

func New() Config {
	dir, err := filepath.Abs("..")
	if err != nil {
		log.Printf("Ошибка пути к файлу настроек:\n%v", err)
	}
	pathConfigFile := filepath.Join(dir, "configs", "config.json")
	configByte, err := os.ReadFile(pathConfigFile)
	if err != nil {
		log.Printf("Ошибка получения настроек:\n%v", err)
		os.Exit(1)
	}
	var config Config
	if err = json.Unmarshal(configByte, &config); err != nil {
		log.Printf("Ошибка получения настроек:\n%v", err)
		os.Exit(1)
	}
	return config
}
