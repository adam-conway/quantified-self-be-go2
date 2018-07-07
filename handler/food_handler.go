package handler

import (
  "net/http"
  "encoding/json"

  "github.com/gorilla/mux"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/adam-conway/quantified-self-be-go/models"
)

func GetFoods(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  (*w).Header().Set("Access-Control-Allow-Origin", "*")
  foods := []models.Food{}
  db.Find(&foods)
  RespondJSON(w, http.StatusOK, foods)
}

func GetFood(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  food := models.Food{}
  params := mux.Vars(r)

  db.First(&food, params["id"])
  RespondJSON(w, http.StatusOK, food)
}

func DeleteFood(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  food := models.Food{}
  params := mux.Vars(r)

  db.First(&food, params["id"])
  db.Delete(&food)

  RespondJSON(w, http.StatusOK, food)
}

func UpdateFood(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  food := models.Food{}
  params := mux.Vars(r)
  db.First(&food, params["id"])
  err := json.NewDecoder(r.Body).Decode(&food)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
  db.Save(&food)

  RespondJSON(w, http.StatusOK, food)
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
