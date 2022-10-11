package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Jaim010/jaim-io/backend/pkg/controllers"
	"github.com/Jaim010/jaim-io/backend/pkg/models"
	"github.com/Jaim010/jaim-io/backend/pkg/utils/testutils"
	"github.com/Jaim010/jaim-io/backend/tests/mocks"
	"github.com/gin-gonic/gin"
)

func TestGetAllExercises(t *testing.T) {
	// Arrange
	expectedExs := []models.Exercise{
		{Id: 1, Name: "Barbell bench press", MuscleGroups: []string{"Chest", "Tricep"}, Description: "Lorem ipsum", ImagePath: "/images/bb_bench_press", VideoLink: "https://www.youtube.com/"},
		{Id: 2, Name: "Bulgarian split squat", MuscleGroups: []string{"Quad", "Glute"}, Description: "Lorem ipsum", ImagePath: "/images/b_split_squad", VideoLink: "https://www.youtube.com/"},
	}

	env := controllers.Env{ExerciseContext: &mocks.MockExerciseContext{}}
	router := gin.Default()
	router.GET("/exercise/", env.GetAllExercises)

	req, _ := http.NewRequest("GET", "/exercise/", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	testutils.CheckResponseCode(t, http.StatusOK, w.Code)

	var responeExs []models.Exercise
	json.Unmarshal(w.Body.Bytes(), &responeExs)
	testutils.CheckEqual(t, expectedExs, responeExs)
}

type getIdExerciseTest struct {
	GivenId          string
	ExpectedExercise models.Exercise
	ExpectedCode     int
}

var getIdExerciseTests = []getIdExerciseTest{
	{
		GivenId:          "1",
		ExpectedExercise: models.Exercise{Id: 1, Name: "Barbell bench press", MuscleGroups: []string{"Chest", "Tricep"}, Description: "Lorem ipsum", ImagePath: "/images/bb_bench_press", VideoLink: "https://www.youtube.com/"},
		ExpectedCode:     200},
	{
		GivenId:          "-1",
		ExpectedExercise: models.Exercise{},
		ExpectedCode:     400},
	{
		GivenId:          "a",
		ExpectedExercise: models.Exercise{},
		ExpectedCode:     400},
	{
		GivenId:          "1000",
		ExpectedExercise: models.Exercise{},
		ExpectedCode:     404},
}

func TestGetExercise(t *testing.T) {
	// Arrange
	env := controllers.Env{ExerciseContext: &mocks.MockExerciseContext{}}
	router := gin.Default()
	router.GET("/exercise/:id", env.GetExerciseById)

	for _, test := range getIdExerciseTests {
		req, _ := http.NewRequest("GET", "/exercise/"+test.GivenId, nil)
		w := httptest.NewRecorder()

		// Act
		router.ServeHTTP(w, req)

		// Assert
		testutils.CheckResponseCode(t, test.ExpectedCode, w.Code)
		if test.ExpectedCode == 200 {
			var responseExercise models.Exercise
			json.Unmarshal(w.Body.Bytes(), &responseExercise)
			testutils.CheckEqual(t, test.ExpectedExercise, responseExercise)
		}
	}
}

type putExerciseTest struct {
	GivenId           string
	GivenExerciseData string
	ExpectedCode      int
}

var putExerciseTests = []putExerciseTest{
	{
		GivenId:           "1",
		GivenExerciseData: `{"id": 1, "name": "Military press", "muscle_groups": ["Shoulder"], "description": "Lorem ipsum", "image_path": "/images/military_press", "video_link": "https://www.youtube.com/"}`,
		ExpectedCode:      204},
	{
		GivenId:           "2",
		GivenExerciseData: `{"id": 1, "name": "Military press", "muscle_groups": ["Shoulder"], "description": "Lorem ipsum", "image_path": "/images/military_press", "video_link": "https://www.youtube.com/"}`,
		ExpectedCode:      400},
	{
		GivenId:           "a",
		GivenExerciseData: `{"id": 1, "name": "Military press", "muscle_groups": ["Shoulder"], "description": "Lorem ipsum", "image_path": "/images/military_press", "video_link": "https://www.youtube.com/"}`,
		ExpectedCode:      400},
	{
		GivenId:           "3",
		GivenExerciseData: `{"id": 3, "name": "Military press", "muscle_groups": ["Shoulder"], "description": "Lorem ipsum", "image_path": "/images/military_press", "video_link": "https://www.youtube.com/"}`,
		ExpectedCode:      404},
}

func TestPutExercise(t *testing.T) {
	// Arrange
	env := controllers.Env{ExerciseContext: &mocks.MockExerciseContext{}}
	router := gin.Default()
	router.PUT("/exercise/:id", env.PutExercise)

	for _, test := range putExerciseTests {
		jsonStrExercise := []byte(test.GivenExerciseData)
		req, _ := http.NewRequest("PUT", "/exercise/"+test.GivenId, bytes.NewBuffer(jsonStrExercise))
		w := httptest.NewRecorder()

		// Act
		router.ServeHTTP(w, req)

		// Assert
		testutils.CheckResponseCode(t, test.ExpectedCode, w.Code)
	}
}

type postExerciseTest struct {
	GivenExerciseData string
	ExpectedCode      int
}

var postExerciseTests = []postExerciseTest{
	{
		GivenExerciseData: `{"name": "Military press", "muscle_groups": ["Shoulder"], "description": "Lorem ipsum", "image_path": "/images/military_press", "video_link": "https://www.youtube.com/"}`,
		ExpectedCode:      201},
	{
		GivenExerciseData: `{"worked-muscles": "tricep"}`,
		ExpectedCode:      400},
}

func TestPostExercise(t *testing.T) {
	// Arrange
	env := controllers.Env{ExerciseContext: &mocks.MockExerciseContext{}}
	router := gin.Default()
	router.POST("/exercise/", env.PostExercise)

	for _, test := range postExerciseTests {
		jsonStrExercise := []byte(test.GivenExerciseData)
		req, _ := http.NewRequest("POST", "/exercise/", bytes.NewBuffer(jsonStrExercise))
		w := httptest.NewRecorder()

		// Act
		router.ServeHTTP(w, req)

		// Assert
		testutils.CheckResponseCode(t, test.ExpectedCode, w.Code)
	}
}

type deleteExerciseTest struct {
	GivenId      string
	ExpectedCode int
}

var deleteExerciseTests = []deleteExerciseTest{
	{
		GivenId:      "1",
		ExpectedCode: 204},
	{
		GivenId:      "a",
		ExpectedCode: 400},
	{
		GivenId:      "1000",
		ExpectedCode: 404},
}

func TestDeleteExercise(t *testing.T) {
	// Arrange
	env := controllers.Env{ExerciseContext: &mocks.MockExerciseContext{}}
	router := gin.Default()
	router.DELETE("/exercise/:id", env.DeleteExercise)

	for _, test := range deleteExerciseTests {
		req, _ := http.NewRequest("DELETE", "/exercise/"+test.GivenId, nil)
		w := httptest.NewRecorder()

		// Act
		router.ServeHTTP(w, req)

		// Assert
		testutils.CheckResponseCode(t, test.ExpectedCode, w.Code)
	}
}
