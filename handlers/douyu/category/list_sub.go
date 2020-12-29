package category

import (
	"net/http"

	"github.com/0987363/agent/middleware"
	"github.com/0987363/agent/models"
	"github.com/gin-gonic/gin"
)

func ListSub(c *gin.Context) {
	logger := middleware.GetLogger(c)
	//	db := middleware.GetBolt(c)

	shortName := c.Param("short_name")
	if shortName == "" {
		logger.Error("Recved unknown short name.")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	dcs, err := models.ListDouyuSubCategory(shortName)
	if err != nil {
		logger.Error("List douyu sub category failed:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, dcs)
}
