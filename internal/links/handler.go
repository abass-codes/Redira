package links

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

type CreateLinkRequest struct {
	URL string `json:"url" binding:"required,url"`
}

type CreateLinkResponse struct {
	ID          string `json:"id"`
	OriginalURL string `json:"originalUrl"`
	ShortCode   string `json:"shortCode"`
	ShortURL    string `json:"shortUrl"`
}

func (h *Handler) Create(c *gin.Context) {
	var req CreateLinkRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	link, err := h.service.Create(c.Request.Context(), req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, CreateLinkResponse{
		ID:          link.ID.String(),
		OriginalURL: link.OriginalUrl,
		ShortCode:   link.ShortCode,
		ShortURL:    "http://localhost:8080/r/" + link.ShortCode,
	})
}
