package config

import (
	"github.com/spf13/viper"
	"log/slog"
)

func LoadConfig() (*Config, error) {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	cnf := getDefaultConfig()

	if err := viper.Unmarshal(cnf); err != nil {
		slog.Error("failed to unmarshal config", "error", err.Error())
		return nil, err
	}

	return cnf, nil
}

type Config struct {
	ListenIP   string
	ListenPort int

	SignalingHost string

	StunServers []string
}

func getDefaultConfig() *Config {

	return &Config{
		ListenIP:      "127.0.0.1",
		ListenPort:    8080,
		SignalingHost: "ws://127.0.0.1:38080/ws",
		StunServers:   []string{"stun:stun.l.google.com:19302"},
	}
}
