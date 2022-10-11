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

func TestGetAllMuscleGroups(t *testing.T) {
	// Arrange
	expectedMgs := []models.MuscleGroup{
		{Id: 1, Name: "Chest", Description: "Lorem ipsum", ImagePath: "/images/chest"},
		{Id: 2, Name: "Tricep", Description: "Lorem ipsum", ImagePath: "/images/tricep"},
	}

	env := controllers.Env{MuscleGroupContext: &mocks.MockMuscleGroupContext{}}
	router := gin.Default()
	router.GET("/musclegroup/", env.GetAllMuscleGroups)

	req, _ := http.NewRequest("GET", "/musclegroup/", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	testutils.CheckResponseCode(t, http.StatusOK, w.Code)

	var responeMgs []models.MuscleGroup
	json.Unmarshal(w.Body.Bytes(), &responeMgs)
	testutils.CheckEqual(t, expectedMgs, responeMgs)
}

type getIdMuscleGroupTest struct {
	GivenId             string
	ExpectedMuscleGroup models.MuscleGroup
	ExpectedCode        int
}

var getIdMuscleGroupTests = []getIdMuscleGroupTest{
	{
		GivenId:             "1",
		ExpectedMuscleGroup: models.MuscleGroup{Id: 1, Name: "Chest", Description: "Lorem ipsum", ImagePath: "/images/chest"},
		ExpectedCode:        200},
	{
		GivenId:             "-1",
		ExpectedMuscleGroup: models.MuscleGroup{},
		ExpectedCode:        400},
	{
		GivenId:             "a",
		ExpectedMuscleGroup: models.MuscleGroup{},
		ExpectedCode:        400},
	{
		GivenId:             "1000",
		ExpectedMuscleGroup: models.MuscleGroup{},
		ExpectedCode:        404},
}

func TestGetMuscleGroup(t *testing.T) {
	// Arrange
	env := controllers.Env{MuscleGroupContext: &mocks.MockMuscleGroupContext{}}
	router := gin.Default()
	router.GET("/musclegroup/:id", env.GetMuscleGroupById)

	for _, test := range getIdMuscleGroupTests {
		req, _ := http.NewRequest("GET", "/musclegroup/"+test.GivenId, nil)
		w := httptest.NewRecorder()

		// Act
		router.ServeHTTP(w, req)

		// Assert
		testutils.CheckResponseCode(t, test.ExpectedCode, w.Code)
		if test.ExpectedCode == 200 {
			var responseMuscleGroup models.MuscleGroup
			json.Unmarshal(w.Body.Bytes(), &responseMuscleGroup)
			testutils.CheckEqual(t, test.ExpectedMuscleGroup, responseMuscleGroup)
		}
	}
}

type putMuscleGroupTest struct {
	GivenId              string
	GivenMuscleGroupData string
	ExpectedCode         int
}

var putMuscleGroupTests = []putMuscleGroupTest{
	{
		GivenId:              "1",
		GivenMuscleGroupData: `{"id": 1, "name": "Back", "description": "Lorem ipsum", "image_path": "/images/back"}`,
		ExpectedCode:         204},
	{
		GivenId:              "2",
		GivenMuscleGroupData: `{"id": 1, "name": "Back", "description": "Lorem ipsum", "image_path": "/images/back"}`,
		ExpectedCode:         400},
	{
		GivenId:              "a",
		GivenMuscleGroupData: `{"id": 1, "name": "Back", "description": "Lorem ipsum", "image_path": "/images/back"}`,
		ExpectedCode:         400},
	{
		GivenId:              "3",
		GivenMuscleGroupData: `{"id": 3, "name": "Back", "description": "Lorem ipsum", "image_path": "/images/back"}`,
		ExpectedCode:         404},
}

func TestPutMuscleGroup(t *testing.T) {
	// Arrange
	env := controllers.Env{MuscleGroupContext: &mocks.MockMuscleGroupContext{}}
	router := gin.Default()
	router.PUT("/musclegroup/:id", env.PutMuscleGroup)

	for _, test := range putMuscleGroupTests {
		jsonStrMuscleGroup := []byte(test.GivenMuscleGroupData)
		req, _ := http.NewRequest("PUT", "/musclegroup/"+test.GivenId, bytes.NewBuffer(jsonStrMuscleGroup))
		w := httptest.NewRecorder()

		// Act
		router.ServeHTTP(w, req)

		// Assert
		testutils.CheckResponseCode(t, test.ExpectedCode, w.Code)
	}
}

type postMuscleGroupTest struct {
	GivenMuscleGroupData string
	ExpectedCode         int
}

var postMuscleGroupTests = []postMuscleGroupTest{
	{
		GivenMuscleGroupData: `{"name": "Back", "description": "Lorem ipsum", "image_path": "/images/back"}`,
		ExpectedCode:         201},
	{
		GivenMuscleGroupData: `{"animal": "cat"}`,
		ExpectedCode:         400},
}

func TestPostMuscleGroup(t *testing.T) {
	// Arrange
	env := controllers.Env{MuscleGroupContext: &mocks.MockMuscleGroupContext{}}
	router := gin.Default()
	router.POST("/musclegroup/", env.PostMuscleGroup)

	for _, test := range postMuscleGroupTests {
		jsonStrMuscleGroup := []byte(test.GivenMuscleGroupData)
		req, _ := http.NewRequest("POST", "/musclegroup/", bytes.NewBuffer(jsonStrMuscleGroup))
		w := httptest.NewRecorder()

		// Act
		router.ServeHTTP(w, req)

		// Assert
		testutils.CheckResponseCode(t, test.ExpectedCode, w.Code)
	}
}

type deleteMuscleGroupTest struct {
	GivenId      string
	ExpectedCode int
}

var deleteMuscleGroupTests = []deleteMuscleGroupTest{
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

func TestDeleteMuscleGroup(t *testing.T) {
	// Arrange
	env := controllers.Env{MuscleGroupContext: &mocks.MockMuscleGroupContext{}}
	router := gin.Default()
	router.DELETE("/musclegroup/:id", env.DeleteMuscleGroup)

	for _, test := range deleteMuscleGroupTests {
		req, _ := http.NewRequest("DELETE", "/musclegroup/"+test.GivenId, nil)
		w := httptest.NewRecorder()

		// Act
		router.ServeHTTP(w, req)

		// Assert
		testutils.CheckResponseCode(t, test.ExpectedCode, w.Code)
	}
}
