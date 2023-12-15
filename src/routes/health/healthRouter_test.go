package healthRouter

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	router = gin.Default()
	NewHealthRouter().Init(router.Group("/"))
	os.Exit(m.Run())
}

func fireGet(path string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", path, nil)
	router.ServeHTTP(w, req)
	return w
}

func TestPingRoute(t *testing.T) {
	w := fireGet("/ping")
	assert.Equal(t, 200, w.Code)
}

func TestInfoRoute(t *testing.T) {
	w := fireGet("/info")
	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"state":"UP","version":"master"}`, w.Body.String())
}

func TestHealthRoute(t *testing.T) {
	w := fireGet("/health")
	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"status":"OK"}`, w.Body.String())
}
