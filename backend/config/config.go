package config

import (
    "log"
    "github.com/spf13/viper"
)

func Load() {
    viper.SetConfigFile(".env")

    if err:= viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading .env file: %v", err)
    }

    viper.SetDefault("APP_PORT", "8080")
}
