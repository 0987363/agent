package room

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/0987363/agent/middleware"
	"github.com/gin-gonic/gin"
)

func GetV3(c *gin.Context) {
	logger := middleware.GetLogger(c)

	roomID := c.Param("id")
	if roomID == "" {
		logger.Error("Recved unknown room id.")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var cmd *exec.Cmd
	if _, err := os.Stat("douyu-full.py"); err == nil {
		cmd = exec.Command("python3", "douyu-full.py", roomID)
	} else if _, err := os.Stat("py/douyu-full.py"); err == nil {
		cmd = exec.Command("python3", "py/douyu-full.py", roomID)
	} else if _, err := os.Stat("/douyu-full.py"); err == nil {
		cmd = exec.Command("python3", "/douyu-full.py", roomID)
	} else {
		logger.Error("Could not found python file.")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//#EXTINF:-1 tvg-logo="null" group-title="%s", test3
	//http://douyu.home.coolhei.com:9009/v2/douyu/room/id/6554039

	out, err := cmd.CombinedOutput()
	if err != nil {
		logger.Errorf("Load media source from node failed: %v, response: %v", err, out)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	res := strings.ReplaceAll(string(out), "'", "\"")

	m := make(map[string]string)
	if err := json.Unmarshal([]byte(res), &m); err != nil {
		logger.Errorf("Json unmarshal failed: %v, response: %v", err, string(out))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	logger.Info("Response: ", m)
	result := "#EXTM3U\n"
	for k, v := range m {
		result = fmt.Sprintf("%s#EXTINF:-1 tvg-logo=null group-title=%s, %s\n", result, roomID, k)
		result = fmt.Sprintf("%s%s\n", result, v)
	}
	logger.Info("Result: ", result)

	c.String(http.StatusOK, result)
}
