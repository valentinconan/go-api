package api



import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestApiRouteOK(t *testing.T) {
	router := gin.Default()
	Init(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/123", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestApiRouteKO(t *testing.T) {
	router := gin.Default()
	Init(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/THIS-PARAMETER-MUST-BE-AN-INTEGER", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}