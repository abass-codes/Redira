package redirect

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(
	service *Service,
) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Redirect(
	c *gin.Context,
) {
	shortCode := c.Param("shortCode")

	log.Println("REDIRECT SHORT CODE:", shortCode)

	url, err := h.service.Redirect(
		c.Request.Context(),
		shortCode,
	)

	if err != nil {
		log.Println("REDIRECT FAILED:", err)

		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	log.Println("REDIRECT SUCCESS:", url)

	c.Redirect(
		http.StatusFound,
		url,
	)
}
