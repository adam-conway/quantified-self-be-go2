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
      
      // "dbname=d4r0mtqvg964h6
      // host=ec2-107-22-174-187.compute-1.amazonaws.com
      // port=5432
      // user=ibzmywfflhxzuc
      // password=8c42e1e51a2bf5660a9d104bb726067eb375b484bd367e585a3c447ee88f1971
      // sslmode=require"
    },
  }
}
