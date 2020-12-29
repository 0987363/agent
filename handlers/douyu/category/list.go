package category

import (
	"net/http"

	"github.com/0987363/agent/middleware"
	"github.com/0987363/agent/models"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	logger := middleware.GetLogger(c)
	//	db := middleware.GetBolt(c)

	dcs, err := models.ListDouyuCategory()
	if err != nil {
		logger.Error("List douyu category failed:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, dcs)
}
