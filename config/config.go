package config

import (
  "os"
)

type DBConfig struct {
  Host string
  Port string
  User string
  DBName string
  SSLMode string
}

type Config struct {
  DB *DBConfig
}

func GetConfig() *Config {
  return &Config {
    DB: &DBConfig {
      Host: "heroku-postgresql",
      Port: "5432",
      // User: "adamconway",
      User: os.Getenv("QS_GO_DB_USERNAME"),
      // DBName: "adamconway",
      DBName: os.Getenv("DATABASE_URL"),
      SSLMode: "disable",
    },
  }
}
