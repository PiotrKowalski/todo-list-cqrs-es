package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

var (
	ErrNotSet   = errors.New("environment variable not set")
	ErrNotValid = errors.New("environment variable not valid")
)

func ReadEnvInt(key string) (int, error) {
	ok := viper.IsSet(key)
	if !ok {
		return 0, fmt.Errorf("key=%v;  %w", key, ErrNotSet)
	}

	value := viper.GetInt(key)
	if value == 0 {
		return 0, fmt.Errorf("key=%v; %w", key, ErrNotValid)
	}

	return value, nil
}

func ReadEnvString(key string) (string, error) {
	ok := viper.IsSet(key)
	if !ok {
		return "", fmt.Errorf("key=%v;  %w", key, ErrNotSet)
	}

	value := viper.GetString(key)
	if value == "" {
		return "", fmt.Errorf("key=%v; %w", key, ErrNotValid)
	}

	return value, nil
}

func init() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
}
