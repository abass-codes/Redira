package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.New()

	router.GET(
		"/health",
		func(c *gin.Context) {
			c.JSON(
				http.StatusOK,
				gin.H{
					"status": "ok",
				},
			)
		},
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/health",
		nil,
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(
		recorder,
		req,
	)

	assert.Equal(
		t,
		http.StatusOK,
		recorder.Code,
	)

	assert.JSONEq(
		t,
		`{"status":"ok"}`,
		recorder.Body.String(),
	)
}
