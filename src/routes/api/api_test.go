package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine
func TestMain(m *testing.M) {
	router = gin.Default()
	Init(router.Group("/api"))
    os.Exit(m.Run())
}

func fireGet(path string) *httptest.ResponseRecorder{
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", path, nil)
	router.ServeHTTP(w, req)
	return w
}

func TestApiRouteOK(t *testing.T) {
	w := fireGet("/api/123")
	assert.Equal(t, 200, w.Code)
}

func TestApiRouteKO(t *testing.T) {
	w := fireGet("/api/THIS-PARAMETER-MUST-BE-AN-INTEGER")
	assert.Equal(t, 400, w.Code)
}