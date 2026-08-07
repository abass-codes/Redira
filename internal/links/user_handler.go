package links

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	repository *UserRepository
}

func NewUserHandler(
	repository *UserRepository,
) *UserHandler {

	return &UserHandler{
		repository: repository,
	}
}

func (h *UserHandler) GetMyLinks(
	c *gin.Context,
) {

	userID := c.MustGet("user_id").(uuid.UUID)

	links, err := h.repository.GetUserLinks(
		c.Request.Context(),
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
		http.StatusOK,
		gin.H{
			"links": links,
		},
	)
}

func (h *UserHandler) DeleteMyLink(
	c *gin.Context,
) {

	userID := c.MustGet("user_id").(uuid.UUID)

	linkID, err := uuid.Parse(
		c.Param("id"),
	)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid id",
			},
		)

		return
	}

	err = h.repository.DeleteUserLink(
		c.Request.Context(),
		linkID,
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
		http.StatusOK,
		gin.H{
			"message": "deleted",
		},
	)
}
