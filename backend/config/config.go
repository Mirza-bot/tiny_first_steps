package config

import (
    "context"
    "log"
    "fmt"
    "github.com/jackc/pgx/v5"
    "github.com/spf13/viper"
)

var DB *pgx.Conn

func Load() {
    viper.SetConfigFile(".env")

    if err:= viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading .env file: %v", err)
    }

    viper.SetDefault("APP_PORT", "8080")
}

func ConnectDatabase() {
    var err error
    dbURL := fmt.Sprintf("postgres://%s:%s@postgres:5432/%s?sslmode=disable",
    viper.GetString("POSTGRES_USER"),
    viper.GetString("POSTGRES_PASSWORD"),
    viper.GetString("POSTGRES_DB"),
    )

    DB, err = pgx.Connect(context.Background(), dbURL)
    if err != nil {
        log.Fatalf("Unable to connect to database %v \n", err)
    }

    fmt.Println("Connected to PostgreSQL!")
}
