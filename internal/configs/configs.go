package configs

import (
	"github.com/CyclopsV/check-sub-tgbot/pkg/pars"
)

type Configs struct {
	BotToken string   `json:"bot_token"`
	DB       DBConfig `json:"database"`
}

func New(path string) (Configs, error) {
	c := Configs{}
	if err := pars.JSON(&c, path); err != nil {
		return c, err
	}
	return c, nil
}
