package room

import (
	"net/http"

	"github.com/0987363/agent/middleware"
	"github.com/0987363/agent/models"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	logger := middleware.GetLogger(c)
	//	db := middleware.GetBolt(c)

	page := models.DecodePage(c)
	dcs, err := models.ListDouyuRoom(page)
	if err != nil {
		logger.Error("List douyu room failed:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, dcs)
}
