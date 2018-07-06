package handler

import (
  "net/http"
  "fmt"
  "strconv"

  "github.com/gorilla/mux"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/adam-conway/quantified-self-be-go/models"
)

func CreateMealFood(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  meal := models.Meal{}
  food := models.Food{}
  food_id, err := strconv.ParseUint(params["id"], 10, 32)
  if err != nil {
        fmt.Println(err)
    }
  foodID := uint(food_id)
  
  meal_id, err := strconv.ParseUint(params["meal_id"], 10, 32)
  if err != nil {
        fmt.Println(err)
    }
  mealID := uint(meal_id)

  db.First(&food, params["id"])
  db.First(&meal, params["meal_id"])
  mealFood := &models.MealFood{MealID: mealID, FoodID: foodID}
  db.NewRecord(mealFood)
  if err := db.Create(&mealFood).Error; err != nil {
    RespondError(w, http.StatusInternalServerError, err.Error())
  } else {
    message := map[string]string{"message": "Successfully added " + food.Name + " to " + meal.Name}
    RespondJSON(w, http.StatusCreated, message)
  }
}
