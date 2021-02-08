## Go Gym Rat Application

An API for keeping a track of workouts at the gym.

#### Start up
Run the database and server with `docker-compose up -d`  

Api is also available at:  
http://ec2co-ecsel-h2ik0akzd8q8-979290349.us-west-2.elb.amazonaws.com/

### DB Relations

WORKOUT * ---> * Exercise 1 ---> * Day 1 ---> * ExerciseSets

### API

A Workout represents a group of exercises at the gym.  
A Workout consists of:  

{  
  id,  
  name  
}  

#### Get all Workouts  
GET /api/workouts  

#### Create a new Workout  
POST /api/workout  
###### BodyParams
  name - Name of the Workout  

#### Update a workout  
PUT /api/workout/:id  
##### UrlParams  
  id - Id of the Workout to update  
###### BodyParams
  name - New name of the Workout  

#### Delete a workout  
DELETE /api/workout/:id  
##### UrlParams  
  id - Id of the workout to update  

-------------------------------------------------------------------------------  

An Exercise represents an exercise at the gym and can belong to many Workouts.  
An Exercise consists of:  

{  
  id,  
  name  
}  

#### Get an Exercise  
GET /api/exercise/:id  
###### UrlParams  
  id - id of the exercise to fetch  

#### Get all Exercises  
GET /api/exercises  

#### Create a new Exercise  
POST /api/exercise  
###### BodyParams
  name - Name of the Exercise  

#### Update an Exercise  
PUT api/exercise/:id  
###### UrlParams  
  id - Id of the Exercise to update  
###### BodyParams
  name - New name of the Exercise  

#### Delete an Exercise  
DELETE /api/exercise/:id  
###### UrlParams  
  id - Id of the Workout to update  

-------------------------------------------------------------------------------  

A WorkoutExercise represents the many-many relation ship between Workouts and  
Exercises.  
A WorkoutExercise consists of:  

{  
  workoutId,  
  exerciseId  
}  

#### Gets all Exercises associated with a Workout  
GET /api/workoutexercise/:id  
###### UrlParams  
  id - Id of the Workout to fetch exercises from  

#### Creates an association between a Workout and an Exercise  
POST /api/workoutexercise  
###### BodyParams
  workoutId - Id of the Workout  
  exerciseId - Id of the Exercise  

#### Deletes an association between a Workout and an Exercise  
DELETE /api/workout/:id/exercise/:exerciseId"  
###### UrlParams  
  id - Id of the Workout to remove the exercise from  
  exerciseId - Id of the exercise to remove from the workout  

-------------------------------------------------------------------------------  

A Day represents a group of sets at the gym.  
A Day consists of  

{  
  id,  
  dayDate,  
  exerciseId  
}  

#### Get Day by id  
GET /api/day/:id  
###### UrlParams  
  id - Id of the Day to get  

#### Get all days related to an exercise  
GET /api/days/:id  
###### UrlParams  
  id - Id of the Exercise to get all days for  

#### Create a new Day  
POST /api/day  
####### BodyParams
  exerciseId - Id of the Exercise to create a new Day for  

#### Delete a Day  
DELETE /api/day/:id  
###### UrlParams  
  id - Id of the Day to delete  

-------------------------------------------------------------------------------  

An ExerciseSet represents a set at the gym. An Exercise set will have the weight  
used and the the number of reps done. SetNo represents which number set it is  
and as you add more sets to the day this increases sequentially.  

An ExerciseSet consists of  
{  
  id,  
  dayId,  
  setNo,  
  rep,  
  weight  
}  

#### Get an ExerciseSet by ID  
GET /api/exerciseset/:id  
###### UrlParams  
  id - Id of the ExerciseSet to get all  

#### Get all ExerciseSets from a day  
GET /api/exercisesets/:id  
###### UrlParams  
  id - Id of the Day to get ExerciseSets from  

#### Create a new ExerciseSet  
POST /api/exerciseset  
###### BodyParams
  dayId - Id of the day to add the exercise set to  
  rep - The number of reps done in the set  
  weight - The weight used in the set  

#### Update an Exercise Set  
PUT /api/exerciseset/:id  
###### UrlParams  
  id - Id of the ExerciseSet to update  
###### BodyParams
  dayId - The new day the set is apart of  
  rep - The new number of reps  
  weight - The new weight used  

#### Delete an ExerciseSet  
DELETE /api/exerciseset/:id  
###### UrlParams  
  id - Id of the ExercsiseSet to delete  
