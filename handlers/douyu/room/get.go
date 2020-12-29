package room

import (
	"net/http"
	"os/exec"
	"strings"

	"github.com/0987363/agent/middleware"
	"github.com/0987363/agent/models"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	logger := middleware.GetLogger(c)

	roomID := c.Param("id")
	if roomID == "" {
		logger.Error("Recved unknown room id.")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	cmd := exec.Command("python3", "py/douyu.py", roomID)
	out, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error("Load media source from node failed:", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, &models.DouyuSource{RoomID: roomID, Source: strings.TrimSpace(string(out))})
}
