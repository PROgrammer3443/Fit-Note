package models

type Exercise struct {
    ID        int
    Name      string
    Amount    int
    Unit      string
    Finished  bool
}

type WorkoutPlan struct {
    DailyBurn int
    Exercises []Exercise
}

var WorkoutLibrary = map[int]WorkoutPlan{
    0: {
        DailyBurn: 300,
        Exercises: []Exercise{
            {ID: 1,  Name: "Walking",                Amount: 10, Unit: "minutes", Finished: false},
            {ID: 2,  Name: "Chair Squats",           Amount: 12, Unit: "reps",    Finished: false},
            {ID: 3,  Name: "Wall Push-Ups",          Amount: 15, Unit: "reps",    Finished: false},
            {ID: 4,  Name: "Marching in Place",      Amount:  8, Unit: "minutes", Finished: false},
            {ID: 5,  Name: "Bird Dogs",              Amount: 16, Unit: "reps",    Finished: false},
            {ID: 6,  Name: "Standing Calf Raises",   Amount: 24, Unit: "reps",    Finished: false},
            {ID: 7,  Name: "Glute Bridges",          Amount: 14, Unit: "reps",    Finished: false},
            {ID: 8,  Name: "Arm Circles",            Amount:  8, Unit: "minutes", Finished: false},
            {ID: 9,  Name: "Side Leg Raises",        Amount: 16, Unit: "reps",    Finished: false},
            {ID: 10, Name: "Seated Knee Extensions", Amount: 20, Unit: "reps",    Finished: false},
        },
    },

    1: {
        DailyBurn: 400,
        Exercises: []Exercise{
            {ID: 101, Name: "Brisk Walking",       Amount:  18, Unit: "minutes", Finished: false},
            {ID: 102, Name: "Bodyweight Squats",   Amount:  15, Unit: "reps",    Finished: false},
            {ID: 103, Name: "Push-Ups",            Amount:  10, Unit: "reps",    Finished: false},
            {ID: 104, Name: "Lunges",              Amount:  20, Unit: "reps",    Finished: false},
            {ID: 105, Name: "Plank",               Amount: 120, Unit: "seconds", Finished: false},
            {ID: 106, Name: "Jumping Jacks",       Amount:  60, Unit: "reps",    Finished: false},
            {ID: 107, Name: "Mountain Climbers",   Amount:  40, Unit: "reps",    Finished: false},
            {ID: 108, Name: "Glute Bridges",       Amount:  22, Unit: "reps",    Finished: false},
            {ID: 109, Name: "Resistance Band Rows",Amount:  18, Unit: "reps",    Finished: false},
            {ID: 110, Name: "Bicycle Crunches",    Amount:  30, Unit: "reps",    Finished: false},
        },
    },

    2: {
        DailyBurn: 500,
        Exercises: []Exercise{
            {ID: 201, Name: "Jogging",              Amount:  12, Unit: "minutes", Finished: false},
            {ID: 202, Name: "Goblet Squats",        Amount:  12, Unit: "reps",    Finished: false},
            {ID: 203, Name: "Dumbbell Rows",        Amount:  15, Unit: "reps",    Finished: false},
            {ID: 204, Name: "Push-Ups",             Amount:  12, Unit: "reps",    Finished: false},
            {ID: 205, Name: "Romanian Deadlifts",   Amount:  12, Unit: "reps",    Finished: false},
            {ID: 206, Name: "Shoulder Press",       Amount:  15, Unit: "reps",    Finished: false},
            {ID: 207, Name: "Russian Twists",       Amount:  30, Unit: "reps",    Finished: false},
            {ID: 208, Name: "Step-Ups",             Amount:  30, Unit: "reps",    Finished: false},
            {ID: 209, Name: "Burpees",              Amount:  10, Unit: "reps",    Finished: false},
            {ID: 210, Name: "Plank Variations",     Amount: 150, Unit: "seconds", Finished: false},
        },
    },

    3: {
        DailyBurn: 600,
        Exercises: []Exercise{
            {ID: 301, Name: "Running",            Amount:  10, Unit: "minutes", Finished: false},
            {ID: 302, Name: "Burpees",            Amount:   9, Unit: "reps",    Finished: false},
            {ID: 303, Name: "Jump Squats",        Amount:  18, Unit: "reps",    Finished: false},
            {ID: 304, Name: "Kettlebell Swings",  Amount:  20, Unit: "reps",    Finished: false},
            {ID: 305, Name: "Pull-Ups",           Amount:   8, Unit: "reps",    Finished: false},
            {ID: 306, Name: "Dumbbell Thrusters", Amount:  10, Unit: "reps",    Finished: false},
            {ID: 307, Name: "Box Step-Ups",       Amount:  35, Unit: "reps",    Finished: false},
            {ID: 308, Name: "Push-Ups",           Amount:  18, Unit: "reps",    Finished: false},
            {ID: 309, Name: "Russian Twists",     Amount:  35, Unit: "reps",    Finished: false},
            {ID: 310, Name: "Plank",              Amount: 180, Unit: "seconds", Finished: false},
        },
    },

    4: {
        DailyBurn: 700,
        Exercises: []Exercise{
            {ID: 401, Name: "Sprint Intervals",   Amount:   8, Unit: "minutes", Finished: false},
            {ID: 402, Name: "Burpees",            Amount:   8, Unit: "reps",    Finished: false},
            {ID: 403, Name: "Jump Lunges",        Amount:  20, Unit: "reps",    Finished: false},
            {ID: 404, Name: "Pull-Ups",           Amount:   9, Unit: "reps",    Finished: false},
            {ID: 405, Name: "Kettlebell Swings",  Amount:  22, Unit: "reps",    Finished: false},
            {ID: 406, Name: "Thrusters",          Amount:  12, Unit: "reps",    Finished: false},
            {ID: 407, Name: "Push-Ups",           Amount:  20, Unit: "reps",    Finished: false},
            {ID: 408, Name: "Mountain Climbers",  Amount:  50, Unit: "reps",    Finished: false},
            {ID: 409, Name: "Hanging Leg Raises", Amount:  18, Unit: "reps",    Finished: false},
            {ID: 410, Name: "Plank",              Amount: 210, Unit: "seconds", Finished: false},
        },
    },
}
