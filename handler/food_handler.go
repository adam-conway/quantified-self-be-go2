package handler

import (
  "net/http"
  // "encoding/json"
  // "fmt"

  // "github.com/gorilla/mux"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/adam-conway/quantified-self-be-go/models"
)

func GetFoods(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  foods := []models.Food{}
  db.Find(&foods)
  RespondJSON(w, http.StatusOK, foods)
}
