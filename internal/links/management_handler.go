package links

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ManagementHandler struct {
	service *ManagementService
}

func NewManagementHandler(
	service *ManagementService,
) *ManagementHandler {

	return &ManagementHandler{
		service: service,
	}
}

func (h *ManagementHandler) Get(
	c *gin.Context,
) {

	userID := c.MustGet(
		"user_id",
	).(uuid.UUID)

	id, err := uuid.Parse(
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

	link, err := h.service.Get(
		c.Request.Context(),
		id,
		userID,
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

	c.JSON(
		http.StatusOK,
		link,
	)
}

func (h *ManagementHandler) Update(
	c *gin.Context,
) {

	userID := c.MustGet(
		"user_id",
	).(uuid.UUID)

	id, _ := uuid.Parse(
		c.Param("id"),
	)

	var body struct {
		URL string `json:"url"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {

		c.JSON(
			400,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	link, err := h.service.Update(
		c.Request.Context(),
		id,
		userID,
		body.URL,
	)

	if err != nil {

		c.JSON(
			404,
			gin.H{
				"error": "link not found",
			},
		)

		return
	}

	c.JSON(
		200,
		link,
	)
}

func (h *ManagementHandler) Disable(
	c *gin.Context,
) {

	h.changeStatus(
		c,
		false,
	)
}

func (h *ManagementHandler) Enable(
	c *gin.Context,
) {

	h.changeStatus(
		c,
		true,
	)
}

func (h *ManagementHandler) changeStatus(
	c *gin.Context,
	enable bool,
) {

	userID := c.MustGet(
		"user_id",
	).(uuid.UUID)

	id, _ := uuid.Parse(
		c.Param("id"),
	)

	var (
		link interface{}
		err  error
	)

	if enable {

		link, err = h.service.Enable(
			c.Request.Context(),
			id,
			userID,
		)

	} else {

		link, err = h.service.Disable(
			c.Request.Context(),
			id,
			userID,
		)
	}

	if err != nil {

		c.JSON(
			404,
			gin.H{
				"error": "link not found",
			},
		)

		return
	}

	c.JSON(
		200,
		link,
	)
}
