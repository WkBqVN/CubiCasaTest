package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// provide example to many case to write in short time
func TestGetAll(t *testing.T) {
	hubHandler := HubHandler{}
	hubHandler.Service.InitService()
	r := SetUpRouter()
	t.Run("Case1:(Api is up)", func(t *testing.T) {
		response := `[{"hubId":1,"hubName":"Vietnam             ","hubRegion":"Asia                "},{"hubId":3,"hubName":"Test Hub            ","hubRegion":"Some where          "}]`
		r.GET("/hubs", hubHandler.GetAll)
		req, _ := http.NewRequest("GET", "/hubs", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		responseData, _ := io.ReadAll(w.Body)
		assert.Equal(t, response, string(responseData))
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
func TestGetHubById(t *testing.T) {
	hubHandler := HubHandler{}
	hubHandler.Service.InitService()
	r := SetUpRouter()
	t.Run("Case1:(Id is in Database)", func(t *testing.T) {
		response := `{"hubId":1,"hubName":"Vietnam             ","hubRegion":"Asia                "}`
		r.GET("/hubs/:id", hubHandler.GetHubById)
		req, _ := http.NewRequest("GET", "/hubs/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		responseData, _ := io.ReadAll(w.Body)
		assert.Equal(t, response, string(responseData))
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("Case2:(Id not in Database)", func(t *testing.T) {
		response := `{"Message":"Can't get Hub or hub id is not exits"}`
		req, _ := http.NewRequest("GET", "/hubs/99", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		responseData, _ := io.ReadAll(w.Body)
		assert.Equal(t, response, string(responseData))
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
