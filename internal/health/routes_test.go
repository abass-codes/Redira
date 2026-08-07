package health_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHealth(t *testing.T) {

	router := gin.New()

	router.GET(
		"/health/live",
		func(c *gin.Context) {
			c.JSON(
				200,
				gin.H{
					"status": "alive",
				},
			)
		},
	)

	req :=
		httptest.NewRequest(
			"GET",
			"/health/live",
			nil,
		)

	w :=
		httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

}
