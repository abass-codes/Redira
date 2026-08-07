package links

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

type CreateLinkRequest struct {
	URL string `json:"url" binding:"required"`
}

func (h *Handler) Create(
	c *gin.Context,
) {

	var req CreateLinkRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	var userID *uuid.UUID

	if value, exists := c.Get("user_id"); exists {

		id := value.(uuid.UUID)
		userID = &id
	}

	shortCode := generateShortCode()

	link, err := h.service.Create(
		c.Request.Context(),
		req.URL,
		shortCode,
		userID,
	)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"id":          link.ID,
			"originalUrl": link.OriginalUrl,
			"shortCode":   link.ShortCode,
		},
	)
}

func (h *Handler) Redirect(
	c *gin.Context,
) {

	link, err := h.service.Redirect(
		c.Request.Context(),
		c.Param("shortCode"),
	)

	if err != nil {

		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "link not found",
			},
		)

		return
	}

	c.Redirect(
		http.StatusFound,
		link.OriginalUrl,
	)
}
