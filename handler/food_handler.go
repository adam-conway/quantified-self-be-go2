package handler

import (
  "net/http"
  "encoding/json"
  "fmt"
  "io/ioutil"

  "github.com/gorilla/mux"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/adam-conway/quantified-self-be-go/models"
)

type FoodStruct struct {
    Food models.Food `json:"food"`
}

func GetFoods(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
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
  var food FoodStruct
  body, _ := ioutil.ReadAll(r.Body)
  json.Unmarshal(body, &food)
  db.NewRecord(food.Food)
  if err := db.Create(&food.Food).Error; err != nil {
    RespondError(w, http.StatusInternalServerError, err.Error())
    return
  } else {
    RespondJSON(w, http.StatusCreated, food)
  }
}
