package room

import (
	"net/http"
	"os"
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

	var cmd *exec.Cmd
	if _, err := os.Stat("douyu.py"); err == nil {
		cmd = exec.Command("python3", "douyu.py", roomID)
	} else if _, err := os.Stat("py/douyu.py"); err == nil {
		cmd = exec.Command("python3", "py/douyu.py", roomID)
	} else if _, err := os.Stat("/douyu.py"); err == nil {
		cmd = exec.Command("python3", "/douyu.py", roomID)
	} else {
		logger.Error("Could not found python file.")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		logger.Errorf("Load media source from node failed: %v, response: %v", err, string(out))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, &models.DouyuSource{RoomID: roomID, Source: strings.TrimSpace(string(out))})
}
