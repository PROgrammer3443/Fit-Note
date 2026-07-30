package main

import (
	"FitNote/internal/models"
	"FitNote/internal/validator"
	"log"
	"net/http"
	"strconv"
    "time"
    "strings"

	"golang.org/x/crypto/bcrypt"
)

type LoggedIn struct {
    LoggedIn bool
    User *models.User
    C_Eat int
    C_Burn int

    Form any
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
    userID := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")

    data := LoggedIn{
        LoggedIn: userID != 0,
        User:     nil,
        C_Eat: 0,
        C_Burn: 0,
    }

    AC_Map := map[int]float32{
        0: 1.2,
        1: 1.375,
        2: 1.55,
        3: 1.725,
        4: 1.9,
    }
    
    if userID != 0 {
        if user, err := app.users.Get(userID); err != nil {
            log.Println(err)
        } else {
            data.User = user

            multiplier, ok := AC_Map[user.Ca_level]
            if !ok {
                multiplier = 1.2
            }

            bmr := (10.0 * float64(user.Weight)) +
                (6.25 * float64(user.Height)) -
                (5.0 * float64(user.Age)) +
                5

            maintenance := int(bmr * float64(multiplier))


            Cal_Eaten, err := app.food_logs.GetTotalCaloriesByUser(userID)
            if err != nil {
                log.Println(err)
                return
            }

            Cal_Burned, err := app.workout_logs.GetBurnedByUser(userID)
            if err != nil {
                log.Println(err)
                return
            }

            dailyGoal := 300 + user.Ca_level*100

            remaining := dailyGoal - Cal_Burned
            if remaining < 0 {
                remaining = 0
            }


            data.C_Eat = maintenance - Cal_Eaten
            data.C_Burn = remaining
        }
    }


    app.render(
        w,
        "home.tmpl.html",
        data,
    )
}

type userSignupForm struct {
    Age      int
    Name     string
    Email    string
    Password string
    validator.Validator
}

func (app *application) signup(w http.ResponseWriter, r *http.Request) {
    form := userSignupForm{}

    data := LoggedIn{
        LoggedIn: app.sessionManager.GetInt(r.Context(), "authenticatedUserID") != 0,
        Form: form,
    }

    app.render(
        w,
        "signup.tmpl.html",
        data,
    )
}

func (app *application) signupPost(w http.ResponseWriter, r *http.Request) {
    var form userSignupForm

    age, err := strconv.Atoi(r.FormValue("age"))
    if err != nil {
        form.AddFieldError("age", "Age must be a number")
    }

    form.Age = age
    form.Name = r.FormValue("name")
    form.Email = r.FormValue("email")
    form.Password = r.FormValue("password")

    form.CheckField(validator.ValidAge(form.Age), "age", "Age must be between 10 and 79")
	form.CheckField(validator.ValidUsername(form.Name), "name", "Username must be between 3 and 30 characters")
    form.CheckField(validator.NotBlank(form.Name), "name", "This field cannot be blank")
    form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank")
    form.CheckField(validator.ValidEmail(form.Email), "email", "Enter a valid email address")
    form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")
    form.CheckField(validator.MinChars(form.Password, 8), "password", "Password must be at least 8 characters")

    if !form.Valid() {
        data := LoggedIn{
            LoggedIn: app.sessionManager.GetInt(r.Context(), "authenticatedUserID") != 0,
            Form: form,
        }

        app.render(
            w,
            "signup.tmpl.html",
            data,
        )

        return
    }

    app.sessionManager.Put(
        r.Context(),
        "signupUsername",
        form.Name,
    )

    app.sessionManager.Put(
        r.Context(),
        "signupEmail",
        form.Email,
    )

    app.sessionManager.Put(
        r.Context(),
        "signupPassword",
        form.Password,
    )

    app.sessionManager.Put(
        r.Context(),
        "signupAge",
        strconv.Itoa(form.Age),
    )

    http.Redirect(
        w,
        r,
        "/signup/profile",
        http.StatusSeeOther,
    )
}

type userProfileForm struct {
    Height  int
    Weight  int
    CALevel int
    TALevel int
    validator.Validator
}

func (app *application) profile(w http.ResponseWriter, r *http.Request) {
    form := userProfileForm{}

    data := LoggedIn{
        LoggedIn: app.sessionManager.GetInt(r.Context(), "authenticatedUserID") != 0,
        Form: form,
    }

    app.render(
        w,
        "profile.tmpl.html",
        data,
    )
}

func (app *application) profilePost(w http.ResponseWriter, r *http.Request) {
    var form userProfileForm

    username := app.sessionManager.GetString(
        r.Context(),
        "signupUsername",
    )

    email := app.sessionManager.GetString(
        r.Context(),
        "signupEmail",
    )

    password := app.sessionManager.GetString(
        r.Context(),
        "signupPassword",
    )

    age, err := strconv.Atoi(
        app.sessionManager.GetString(
            r.Context(),
            "signupAge",
        ),
    )
    if err != nil {
        http.Error(w, "Invalid age", http.StatusBadRequest)
        return
    }

    height, err := strconv.Atoi(
        r.FormValue("height"),
    )
    if err != nil {
        form.AddFieldError(
            "height",
            "Height must be a number",
        )
		return
    }

    weight, err := strconv.Atoi(
        r.FormValue("weight"),
    )
    if err != nil {
        form.AddFieldError(
            "weight",
            "Weight must be a number",
        )
		return
    }

    caLevel, err := strconv.Atoi(
        r.FormValue("ca_level"),
    )
    if err != nil {
        form.AddFieldError(
            "ca_level",
            "Choose an activity level",
        )
		return
    }

    taLevel, err := strconv.Atoi(
        r.FormValue("ta_level"),
    )
    if err != nil {
        form.AddFieldError(
            "ta_level",
            "Choose a target activity level",
        )
		return
    }

    form.Height = height
    form.Weight = weight
    form.CALevel = caLevel
    form.TALevel = taLevel

    form.CheckField(validator.ValidHeight(form.Height), "height", "Height must be between 50 cm and 275 cm")
    form.CheckField(validator.ValidWeight(form.Weight), "weight", "Weight must be between 2 kg and 300 kg")

    form.CheckField(
        validator.PermittedValue(
            form.CALevel,
            0, 1, 2, 3, 4,
        ), 
    "ca_level", "Invalid activity level")

    form.CheckField(
        validator.PermittedValue(
            form.TALevel,
            0, 1, 2, 3, 4,
        ),
    "ta_level", "Invalid target activity level")

    if !form.Valid() {
        data := LoggedIn{
            LoggedIn: app.sessionManager.GetInt(r.Context(), "authenticatedUserID") != 0,
            Form: form,
        }

        app.render(
            w,
            "profile.tmpl.html",
            data,
        )

        return
    }

    hash, err := bcrypt.GenerateFromPassword(
        []byte(password),
        bcrypt.DefaultCost,
    )
    if err != nil {
        http.Error(
            w,
            "Server Error",
            http.StatusInternalServerError,
        )
        return
    }

    err = app.users.Insert(
        username,
        email,
        string(hash),
        age,
        form.Height,
        form.Weight,
        form.CALevel,
        form.TALevel,
    )
    if err != nil {
        http.Error(
            w,
            err.Error(),
            http.StatusInternalServerError,
        )
        return
    }

    http.Redirect(
        w,
        r,
        "/login",
        http.StatusSeeOther,
    )
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
    app.render(
        w,
        "login.tmpl.html",
        nil,
    )
}

func (app *application) loginPost(w http.ResponseWriter, r *http.Request) {

    email := r.FormValue("email")
    password := r.FormValue("password")

    user, err := app.users.GetByEmail(email)
    if err != nil {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    err = bcrypt.CompareHashAndPassword(
        []byte(user.Hashed_password),
        []byte(password),
    )

    if err != nil {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    app.sessionManager.Put(
        r.Context(),
        "authenticatedUserID",
        user.ID,
    )

    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) logoutPost(w http.ResponseWriter, r *http.Request) {
	app.sessionManager.Remove(r.Context(), "authenticatedUserID")
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

type FoodPageData struct {
    LoggedIn bool
    Food      []MealView
}

type MealView struct {
    ID              int
    MealType        string
    Date            string
    TotalCalories   int
    Entries         []*models.Logs
}

func (app *application) food(w http.ResponseWriter, r *http.Request) {

    userID := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")

    if userID == 0 {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    foods, err := app.food_logs.GetLogsByUser(userID)
    if err != nil {
        log.Println(err)

        http.Error(
            w,
            "Internal server error",
            http.StatusInternalServerError,
        )

        return
    }

    meals := []MealView{}

    for _, food := range foods {
        entries, err := app.log_entries.GetByMeal(food.Meal_id)
        if err != nil {
            log.Println(err)

            http.Error(
                w,
                "Internal server error",
                http.StatusInternalServerError,
            )

            return
        }

        totalCalories := 0

        for _, entry := range entries {
            totalCalories += entry.Calories
        }

        meals = append(meals, MealView{
            ID:              food.Meal_id,
            MealType:        food.Meal_Type,
            Date:            food.Date.Format("Mon Jan 2"),
            TotalCalories:   totalCalories,
            Entries:         entries,
        })
    }

    data := FoodPageData{
        LoggedIn: true,
        Food:     meals,
    }

    app.render(
        w,
        "food.tmpl.html",
        data,
    )
}

func (app *application) logMealPost(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    userID := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")

    if userID == 0 {
        http.Redirect(
            w,
            r,
            "/login",
            http.StatusSeeOther,
        )
        return
    }

    mealType := r.FormValue("meal_type")

    validMeals := map[string]bool{
        "Breakfast": true,
        "Lunch":     true,
        "Dinner":    true,
        "Snack":     true,
    }

    if !validMeals[mealType] {
        http.Error(
            w,
            "Invalid meal type",
            http.StatusBadRequest,
        )
        return
    }

    err := app.food_logs.Insert(
        mealType,
        userID,
        time.Now())
    if err != nil {
        log.Println(err)

        http.Error(
            w,
            "Internal server error",
            http.StatusInternalServerError,
        )
        return
    }

    http.Redirect(
        w,
        r,
        "/food",
        http.StatusSeeOther,
    )
}

func (app *application) addFoodPost(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    userID := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")

    if userID == 0 {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    mealID, err := strconv.Atoi(
        r.FormValue("meal_id"),
    )
    if err != nil {
        http.Error(w, "Invalid meal", http.StatusBadRequest)
        return
    }

    foodName := strings.TrimSpace(
        r.FormValue("food_name"),
    )

    if foodName == "" {
        http.Error(w, "Food name required", http.StatusBadRequest)
        return
    }

    calories, err := strconv.Atoi(
        r.FormValue("calories"),
    )
    if err != nil || calories < 0 {
        http.Error(w, "Invalid calories", http.StatusBadRequest)
        return
    }

    if err = app.log_entries.Insert(foodName, mealID, calories); err != nil {
        log.Println(err)

        http.Error(
            w,
            "Internal server error",
            http.StatusInternalServerError,
        )

        return
    }

    http.Redirect(
        w,
        r,
        "/food",
        http.StatusSeeOther,
    )
}


type WorkoutPageData struct {
    LoggedIn bool

    DailyBurn int
    ID int

    CurrentExercises []models.Exercise
    Alternatives      []models.Exercise
}

func (app *application) workouts(w http.ResponseWriter, r *http.Request) {
    userID := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")

    if userID == 0 {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    user, err := app.users.Get(userID)
    if err != nil {
        http.Error(w, "Unable to load user", http.StatusInternalServerError)
        return
    }

    plan, ok := models.WorkoutLibrary[user.Ca_level]
    if !ok {
        http.Error(w, "Workout plan not found", http.StatusInternalServerError)
        return
    }

    currentExercises, ok := app.sessionManager.Get(
        r.Context(),
        "currentExercises",
    ).([]models.Exercise)

    if !ok || len(currentExercises) == 0 {
        currentExercises = make([]models.Exercise, 3)
        copy(currentExercises, plan.Exercises[:3])

        app.sessionManager.Put(
            r.Context(),
            "currentExercises",
            currentExercises,
        )
    }

    currentIDs := make(map[int]bool)
    for _, ex := range currentExercises {
        currentIDs[ex.ID] = true
    }

    var alternatives []models.Exercise
    for _, ex := range plan.Exercises {
        if !currentIDs[ex.ID] {
            alternatives = append(alternatives, ex)
        }
    }

    Cal_Burned, err := app.workout_logs.GetBurnedByUser(userID)
    if err != nil {
        log.Println(err)
        return
    }

    remaining := plan.DailyBurn - Cal_Burned
    if remaining < 0 {
        remaining = 0
    }

    plan.DailyBurn = remaining

    data := WorkoutPageData{
        LoggedIn:          true,
        DailyBurn:         plan.DailyBurn,
        CurrentExercises: currentExercises,
        Alternatives:      alternatives,
    }

    app.render(
        w,
        "workout.tmpl.html",
        data,
    )
}

func (app *application) finishWorkoutPost(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    userID := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")

    if userID == 0 {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    exerciseID, err := strconv.Atoi(
        r.FormValue("exercise_id"),
    )
    if err != nil {
        http.Error(w, "Invalid exercise", http.StatusBadRequest)
        return
    }

    currentExercises, ok := app.sessionManager.Get(
        r.Context(),
        "currentExercises",
    ).([]models.Exercise)

    if !ok {
        http.Error(
            w,
            "Workout session not found",
            http.StatusBadRequest,
        )
        return
    }

    found := false

    for i := range currentExercises {

        if currentExercises[i].ID == exerciseID {

            wasFinished := currentExercises[i].Finished

            currentExercises[i].Finished =
                !currentExercises[i].Finished

            found = true

            if !wasFinished {

                if err := app.workout_logs.Insert(userID, 100, time.Now()); err != nil {
                    log.Println(err)
                }
            }

            break
        }
    }

    if !found {
        http.Error(
            w,
            "Exercise not found",
            http.StatusBadRequest,
        )
        return
    }

    app.sessionManager.Put(
        r.Context(),
        "currentExercises",
        currentExercises,
    )

    http.Redirect(
        w,
        r,
        "/workout",
        http.StatusSeeOther,
    )
}

func (app *application) swapWorkoutPost(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    userID := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")

    if userID == 0 {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    currentID, err := strconv.Atoi(
        r.FormValue("current_exercise_id"),
    )
    if err != nil {
        http.Error(w, "Invalid current exercise", http.StatusBadRequest)
        return
    }

    replacementID, err := strconv.Atoi(
        r.FormValue("replacement_exercise_id"),
    )
    if err != nil {
        http.Error(w, "Invalid replacement exercise", http.StatusBadRequest)
        return
    }

    user, err := app.users.Get(userID)
    if err != nil {
        http.Error(w, "Unable to load user", http.StatusInternalServerError)
        return
    }

    plan, ok := models.WorkoutLibrary[user.Ca_level]
    if !ok {
        http.Error(w, "Workout plan not found", http.StatusInternalServerError)
        return
    }

    currentExercises, ok := app.sessionManager.Get(
        r.Context(),
        "currentExercises",
    ).([]models.Exercise)

    if !ok || len(currentExercises) == 0 {
        currentExercises = make([]models.Exercise, 3)
        copy(currentExercises, plan.Exercises[:3])
    }

    var replacement models.Exercise
    foundReplacement := false

    for _, ex := range plan.Exercises {
        if ex.ID == replacementID {
            replacement = ex
            foundReplacement = true
            break
        }
    }

    if !foundReplacement {
        http.Error(w, "Replacement exercise not found", http.StatusBadRequest)
        return
    }

    swapped := false

    for i := range currentExercises {
        if currentExercises[i].ID == currentID {
            currentExercises[i] = replacement
            swapped = true
            break
        }
    }

    if !swapped {
        http.Error(w, "Current exercise not found", http.StatusBadRequest)
        return
    }

    app.sessionManager.Put(
        r.Context(),
        "currentExercises",
        currentExercises,
    )

    http.Redirect(w, r, "/workout", http.StatusSeeOther)
}