package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	"time"
	"encoding/gob"

	_ "github.com/go-sql-driver/mysql"
	"github.com/alexedwards/scs/v2"
	"FitNote/internal/models"
)

type application struct {
	users *models.UserModel
	food_logs *models.FoodModel
    log_entries *models.LogModel
	workout_logs *models.WorkoutModel
	sessionManager *scs.SessionManager
}

func main() {
    gob.Register(models.Exercise{})
    gob.Register([]models.Exercise{})

	db, err := sql.Open(
		"mysql",
		"root:Hamza#901$761@tcp(localhost:3306)/fitnote?parseTime=true",
	)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}

	sessionManager := scs.New()
	sessionManager.Lifetime = 12 * time.Hour
	sessionManager.Cookie.Secure = true

	app := &application{
		users: &models.UserModel{
			DB: db,
		},
		food_logs: &models.FoodModel{
			DB: db,
		},
		log_entries: &models.LogModel{
			DB: db,
		},
		workout_logs: &models.WorkoutModel{
			DB: db,
		},
		sessionManager: sessionManager,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("GET /signup", app.signup)
	mux.HandleFunc("POST /signup", app.signupPost)
	mux.HandleFunc("GET /signup/profile", app.profile)
	mux.HandleFunc("POST /signup/profile", app.profilePost)
	mux.HandleFunc("GET /login", app.login)
	mux.HandleFunc("POST /login", app.loginPost)
	mux.HandleFunc("POST /logout", app.logoutPost)
	mux.HandleFunc("/food", app.food)
	mux.HandleFunc("POST /food/entry/new", app.addFoodPost)
	mux.HandleFunc("POST /food/new", app.logMealPost)
	mux.HandleFunc("/workout", app.workouts)
	mux.HandleFunc("POST /workout/swap", app.swapWorkoutPost)
	mux.HandleFunc("POST /workout/finish", app.finishWorkoutPost)




	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	fmt.Println("Server starting at :4000")

	log.Fatal(
		http.ListenAndServe(
			":4000",
			sessionManager.LoadAndSave(mux),
		),
	)
}