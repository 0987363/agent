package room

import (
	"net/http"

	"github.com/0987363/agent/middleware"
	"github.com/0987363/agent/models"
	"github.com/gin-gonic/gin"
)

func ListCategory(c *gin.Context) {
	logger := middleware.GetLogger(c)
	//	db := middleware.GetBolt(c)

	cateID := c.Param("id")
	if cateID == "" {
		logger.Error("Recved unknown cate id.")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	page := models.DecodePage(c)
	dcs, err := models.ListDouyuRoomByCateID(cateID, page)
	if err != nil {
		logger.Error("List douyu room failed:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, dcs)
}
