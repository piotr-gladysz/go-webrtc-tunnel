package config

import (
	"github.com/spf13/viper"
	"log/slog"
)

func LoadConfig() (*Config, error) {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("relay-config")
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
	Debug bool `mapstructure:"debug"`

	ListenIP   string `mapstructure:"listen_ip"`
	ListenPort int    `mapstructure:"listen_port"`

	DisableTLS    bool   `mapstructure:"disable_tls"`
	SignalingHost string `mapstructure:"signaling_host"`

	StunServers []string `mapstructure:"stun_servers"`
}

func getDefaultConfig() *Config {

	return &Config{
		Debug:         false,
		ListenIP:      "127.0.0.1",
		ListenPort:    13080,
		DisableTLS:    false,
		SignalingHost: "127.0.0.1:38080",
		StunServers:   []string{"stun:stun.l.google.com:19302"},
	}
}
