package models

import (
  // "github.com/jinzhu/gorm"
  // _ "github.com/jinzhu/gorm/dialects/postgres"
)

type MealFood struct {
  // gorm.Model

  ID  uint `json: id`
  Meal Meal
  MealID uint
  Food Food
  FoodID uint
}
