package models

import (
  // "github.com/jinzhu/gorm"
  // _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Food struct {
  // gorm.Model
  ID       uint `json: id`
  Name string `gorm: "type:varchar(100)" json: "name"`
  Calories uint `json: "calories"`
}
