package main

import (
  "net/http"
  "log"
  "fmt"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/gorilla/mux"
  "github.com/adam-conway/quantified-self-be-go/config"
  "github.com/adam-conway/quantified-self-be-go/handler"
)

type App struct {
  DB *gorm.DB
}

func main() {
  config := config.GetConfig()
  a := &App{}
  a.Init(config)
  fmt.Println("Server Running on :8000")

  router := mux.NewRouter()
  router.HandleFunc("/api/v1/foods", a.getFoods).Methods("GET")
  // sub.HandleFunc("/meals", a.GetAllMeals).Methods("GET")
  // sub.HandleFunc("/meals/{id}/foods", a.GetMeal).Methods("GET")
  // sub.HandleFunc("/foods/{id}", a.GetFood).Methods("GET")
  // sub.HandleFunc("/foods", a.CreateFood).Methods("POST")
  // sub.HandleFunc("/meal-foods", a.GetAllMealFoods).Methods("GET")
  // sub.HandleFunc("/meals/{meal_id}/foods/{id}", a.CreateMealFood).Methods("POST")
  // sub.HandleFunc("/meals/{meal_id}/foods/{id}", a.DeleteMealFood).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":8000", router))
}

func (a *App) getFoods(w http.ResponseWriter, r *http.Request) {
  handler.GetFoods(a.DB, w, r)
}

func (a *App) Init(config *config.Config) {
  dbParams := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s",
    config.DB.Host,
    config.DB.Port,
    config.DB.User,
    config.DB.DBName,
    config.DB.SSLMode)

  var err error
  a.DB, err = gorm.Open("postgres", dbParams)
  if err != nil {
    fmt.Println("ERROR IN CONNECTING TO DATABASE\n")
    log.Fatal(err.Error())
  }
  fmt.Println("Connected to Database")

  a.Migrate()
}

func (a *App) Migrate() {
  // a.DB.AutoMigrate(&models.Food{})
  // a.DB.AutoMigrate(&models.Meal{})
  // a.DB.AutoMigrate(&models.MealFood{})
}
