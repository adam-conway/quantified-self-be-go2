package handler

import (
  "net/http"

  // "github.com/gorilla/mux"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/adam-conway/quantified-self-be-go/models"
)

func GetMeals(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  meals := []models.Meal{}
  db.Preload("Foods").Find(&meals)
  RespondJSON(w, http.StatusOK, meals)
}
