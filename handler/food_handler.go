package handler

import (
  "net/http"
  "encoding/json"
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

func CreateFood(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  food := models.Food{}
  decoder := json.NewDecoder(r.Body)
  if err := decoder.Decode(&food); err != nil {
    RespondError(w, http.StatusBadRequest, err.Error())
    return
  }

  defer r.Body.Close()
  db.NewRecord(food)
  if err := db.Create(&food).Error; err != nil {
    RespondError(w, http.StatusInternalServerError, err.Error())
    return
  } else {
    RespondJSON(w, http.StatusCreated, food)
  }
}
