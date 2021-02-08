CREATE TABLE IF NOT EXISTS workout_exercise (
    workoutId VARBINARY(36) NOT NULL,
    exerciseId VARBINARY(36) NOT NULL,
    PRIMARY KEY (workoutId, exerciseId),
    FOREIGN KEY (workoutId) REFERENCES workout(id),
    FOREIGN KEY (exerciseId) REFERENCES exercise(id)
) ENGINE=InnoDB, DEFAULT CHARSET=utf8mb4;
