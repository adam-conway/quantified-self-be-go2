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
      Host: os.Getenv("HOST"),
      Port: os.Getenv("PORT"),
      // User: "adamconway",
      User: os.Getenv("QS_GO_DB_USERNAME"),
      // DBName: "adamconway",
      DBName: os.Getenv("dbname"),
      SSLMode: "disable",
    },
  }
}
