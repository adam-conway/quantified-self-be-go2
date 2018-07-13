package models

import (
  // "github.com/jinzhu/gorm"
  // _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Food struct {
  // gorm.Model
  ID       uint   `json:"id"`
  Name     string `json:"name"`
  Calories string `json:"calories"`
}
