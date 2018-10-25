package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/getndazn/template-go-microservice/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMyApp(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	// Start new writer recorder
	w := httptest.NewRecorder()

	// Setup test router
	r := router.Setup()

	// Create new request
	req, _ := http.NewRequest("GET", "/v1/template/name", nil)

	// Start test router
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
