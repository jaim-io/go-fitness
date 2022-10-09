package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Jaim010/jaim-io/backend/pkg/mocks"
	"github.com/Jaim010/jaim-io/backend/pkg/models"
	"github.com/gin-gonic/gin"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected respone code %d. Got %d\n", expected, actual)
	}
}

func checkEqual[K any](t *testing.T, expected, actual K) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v. Got %v\n", expected, actual)
	}
}

func TestGetAllExercises(t *testing.T) {
	// Arrange
	expectedExs := []models.Exercise{
		{Name: "Barbell bench press", Id: 1, Description: "Lorem ipsum"},
		{Name: "Bulgarian split squat", Id: 2, Description: "Lorem ipsum"},
	}

	env := Env{ExerciseContext: &mocks.MockExerciseContext{}}
	router := gin.Default()
	router.GET("/exercise/", env.GetAllExercises)

	req, _ := http.NewRequest("GET", "/exercise/", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	checkResponseCode(t, http.StatusOK, w.Code)

	var responeExs []models.Exercise
	json.Unmarshal(w.Body.Bytes(), &responeExs)
	checkEqual(t, expectedExs, responeExs)
}

type getIdExerciseTest struct {
	GivenId          string
	ExpectedExercise models.Exercise
	ExpectedCode     int
}

var getIdExerciseTests = []getIdExerciseTest{
	{GivenId: "1", ExpectedExercise: models.Exercise{Id: 1, Name: "Barbell bench press", Description: "Lorem ipsum"}, ExpectedCode: 200},
	{GivenId: "-1", ExpectedExercise: models.Exercise{}, ExpectedCode: 400},
	{GivenId: "a", ExpectedExercise: models.Exercise{}, ExpectedCode: 400},
	{GivenId: "1000", ExpectedExercise: models.Exercise{}, ExpectedCode: 404},
}

func TestGetExercise(t *testing.T) {
	// Arrange
	env := Env{ExerciseContext: &mocks.MockExerciseContext{}}
	router := gin.Default()
	router.GET("/exercise/:id", env.GetExerciseById)

	for _, test := range getIdExerciseTests {
		req, _ := http.NewRequest("GET", "/exercise/"+test.GivenId, nil)
		w := httptest.NewRecorder()

		// Act
		router.ServeHTTP(w, req)

		// Assert
		checkResponseCode(t, test.ExpectedCode, w.Code)
		if test.ExpectedCode == 200 {
			var responseExercise models.Exercise
			json.Unmarshal(w.Body.Bytes(), &responseExercise)
			checkEqual(t, test.ExpectedExercise, responseExercise)
		}
	}
}

type putExerciseTest struct {
	GivenId           string
	GivenExerciseData string
	ExpectedCode      int
}

var putExerciseTests = []putExerciseTest{
	{GivenId: "1", GivenExerciseData: `{"id": 1, "name": "Lunges", "description": "Lorem ipsum"}`, ExpectedCode: 204},
	{GivenId: "2", GivenExerciseData: `{"id": 1, "name": "Lunges", "description": "Lorem ipsum"}`, ExpectedCode: 400},
	{GivenId: "a", GivenExerciseData: `{"id": 1, "name": "Lunges", "description": "Lorem ipsum"}`, ExpectedCode: 400},
	{GivenId: "3", GivenExerciseData: `{"id": 3, "name": "Lunges", "description": "Lorem ipsum"}`, ExpectedCode: 404},
}

func TestPutExercise(t *testing.T) {
	// Arrange
	env := Env{ExerciseContext: &mocks.MockExerciseContext{}}
	router := gin.Default()
	router.PUT("/exercise/:id", env.PutExercise)

	for _, test := range putExerciseTests {
		jsonStrExercise := []byte(test.GivenExerciseData)
		req, _ := http.NewRequest("PUT", "/exercise/"+test.GivenId, bytes.NewBuffer(jsonStrExercise))
		w := httptest.NewRecorder()

		// Act
		router.ServeHTTP(w, req)

		// Assert
		checkResponseCode(t, test.ExpectedCode, w.Code)
	}
}

type postExerciseTest struct {
	GivenExerciseData string
	ExpectedCode      int
}

var postExerciseTests = []postExerciseTest{
	{GivenExerciseData: `{"name": "Tricep extensions", "description": "Lorem ipsum"}`, ExpectedCode: 201},
	{GivenExerciseData: `{"Muscle group": "tricep"}`, ExpectedCode: 400},
}

func TestPostExercise(t *testing.T) {
	// Arrange
	env := Env{ExerciseContext: &mocks.MockExerciseContext{}}
	router := gin.Default()
	router.POST("/exercise/", env.PostExercise)

	for _, test := range postExerciseTests {
		jsonStrExercise := []byte(test.GivenExerciseData)
		req, _ := http.NewRequest("POST", "/exercise/", bytes.NewBuffer(jsonStrExercise))
		w := httptest.NewRecorder()

		// Act
		router.ServeHTTP(w, req)

		// Assert
		checkResponseCode(t, test.ExpectedCode, w.Code)
	}
}

type deleteExerciseTest struct {
	GivenId      string
	ExpectedCode int
}

var deleteExerciseTests = []deleteExerciseTest{
	{GivenId: "1", ExpectedCode: 204},
	{GivenId: "a", ExpectedCode: 400},
	{GivenId: "1000", ExpectedCode: 404},
}

func TestDeleteBook(t *testing.T) {
	// Arrange
	env := Env{ExerciseContext: &mocks.MockExerciseContext{}}
	router := gin.Default()
	router.DELETE("/exercise/:id", env.DeleteExercise)

	for _, test := range deleteExerciseTests {
		req, _ := http.NewRequest("DELETE", "/exercise/"+test.GivenId, nil)
		w := httptest.NewRecorder()

		// Act
		router.ServeHTTP(w, req)

		// Assert
		checkResponseCode(t, test.ExpectedCode, w.Code)
	}
}
