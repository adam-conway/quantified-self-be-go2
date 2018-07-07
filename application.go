package main

import (
  "net/http"
  "log"
  "fmt"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/gorilla/mux"
  "github.com/adam-conway/quantified-self-be-go/models"
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
  router.HandleFunc("/api/v1/foods", a.GetFoods).Methods("GET")
  router.HandleFunc("/api/v1/foods/{id}", a.DeleteFood).Methods("DELETE")
  router.HandleFunc("/api/v1/meals", a.GetMeals).Methods("GET")
  router.HandleFunc("/api/v1/foods/{id}", a.GetFood).Methods("GET")
  router.HandleFunc("/api/v1/foods", a.CreateFood).Methods("POST")
  router.HandleFunc("/api/v1/foods/{id}", a.UpdateFood).Methods("PUT")
  router.HandleFunc("/api/v1/meals/{meal_id}/foods/{id}", a.CreateMealFood).Methods("POST")
  router.HandleFunc("/api/v1/meals/{id}/foods", a.GetMeal).Methods("GET")
  router.HandleFunc("/api/v1/meals/{meal_id}/foods/{id}", a.DeleteMealFood).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":8000", router))
}

func (a *App) GetFoods(w http.ResponseWriter, r *http.Request) {
  handler.GetFoods(a.DB, w, r)
}

func (a *App) GetFood(w http.ResponseWriter, r *http.Request) {
  handler.GetFood(a.DB, w, r)
}

func (a *App) CreateFood(w http.ResponseWriter, r *http.Request) {
  handler.CreateFood(a.DB, w, r)
}

func (a *App) UpdateFood(w http.ResponseWriter, r *http.Request) {
  handler.UpdateFood(a.DB, w, r)
}

func (a *App) DeleteFood(w http.ResponseWriter, r *http.Request) {
  handler.DeleteFood(a.DB, w, r)
}

func (a *App) GetMeals(w http.ResponseWriter, r *http.Request) {
  handler.GetMeals(a.DB, w, r)
}

func (a *App) GetMeal(w http.ResponseWriter, r *http.Request) {
  handler.GetMeal(a.DB, w, r)
}

func (a *App) CreateMealFood(w http.ResponseWriter, r *http.Request) {
  handler.CreateMealFood(a.DB, w, r)
}

func (a *App) DeleteMealFood(w http.ResponseWriter, r *http.Request) {
  handler.DeleteMealFood(a.DB, w, r)
}


func (a *App) Init(config *config.Config) {
  dbParams := fmt.Sprintf("dbname=d4r0mtqvg964h6 host=ec2-107-22-174-187.compute-1.amazonaws.com port=5432 user=ibzmywfflhxzuc password=8c42e1e51a2bf5660a9d104bb726067eb375b484bd367e585a3c447ee88f1971 sslmode=require")
    // config.DB.Host,
    // config.DB.Port,
    // config.DB.User,
    // config.DB.DBName,
    // config.DB.SSLMode)

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
  a.DB.DropTable(&models.Food{})
  a.DB.AutoMigrate(&models.Food{})
  a.DB.DropTable(&models.Meal{})
  a.DB.AutoMigrate(&models.Meal{})
  a.DB.DropTable(&models.MealFood{})
  a.DB.AutoMigrate(&models.MealFood{})

  meal1 := &models.Meal{Name: "breakfast"}
  a.DB.NewRecord(meal1)
  a.DB.Create(&meal1)
  a.DB.NewRecord(meal1)

  meal2 := &models.Meal{Name: "snack"}
  a.DB.NewRecord(meal2)
  a.DB.Create(&meal2)
  a.DB.NewRecord(meal2)

  meal3 := &models.Meal{Name: "lunch"}
  a.DB.NewRecord(meal3)
  a.DB.Create(&meal3)
  a.DB.NewRecord(meal3)

  meal4 := &models.Meal{Name: "dinner"}
  a.DB.NewRecord(meal4)
  a.DB.Create(&meal4)
  a.DB.NewRecord(meal4)
}
