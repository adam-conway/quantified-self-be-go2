package config

import (
  "os"
)

type DBConfig struct {
  Host string
  Port string
  User string
  Password string
  DBName string
  SSLMode string
}

type Config struct {
  DB *DBConfig
}

func GetConfig() *Config {
  return &Config {
    DB: &DBConfig {
      Host: os.Getenv("host"),
      Port: os.Getenv("port"),
      User: os.Getenv("user"),
      Password: os.Getenv("password"),
      DBName: os.Getenv("dbname"),
      SSLMode: os.Getenv("sslmode"),
      // Host: "localhost",
      // Port: "5432",
      // User: "adamconway",
      // DBName: "adamconway",
      // SSLMode: "disable",
    },
  }
}
