package api

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-api/mocks"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	router = gin.Default()
	NewApiRouter().Init(router.Group("/api"))
	os.Exit(m.Run())
}

func fireGet(path string) *httptest.ResponseRecorder {
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

func TestApiRouteRetrieveInfo(t *testing.T) {

	router = gin.Default()

	httpRouter, mockServ := generateRouterMockStruct(t)

	httpRouter.Init(router.Group("/api"))

	mockServ.EXPECT().Call("http://localhost:8080/info").Return("ok")

	w := callGetPath("/api/retrieve/info", "RawBodyTest")

	assert.Equal(t, 200, w.Code)
}

func generateRouterMockStruct(t *testing.T) (*Router, *mocks.MockIHttpService) {
	mockServ := mocks.NewMockIHttpService(t)

	httpRouter := Router{
		Service: mockServ,
	}
	return &httpRouter, mockServ
}

func callGetPath(path string, bodStr string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()

	body := bodStr
	reader := bytes.NewBufferString(body)

	req, _ := http.NewRequest("GET", path, reader)

	router.ServeHTTP(w, req)
	return w
}
