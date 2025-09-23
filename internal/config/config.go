package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Llmserver LLMserver `toml:"llmserver"`
}

type LLMserver struct {
	URL    string `toml:"url"`
	Model  string `toml:"model"`
	Stream bool   `toml:"stream"`
}

var Conf Config

func (conf *Config) Init() error {
	if _, err := toml.DecodeFile("config.toml", conf); err != nil {
		return err
	} else {
		return nil
	}
}
