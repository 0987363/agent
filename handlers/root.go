package handlers

import (
	"time"

	dcate "github.com/0987363/agent/handlers/douyu/category"
	droom "github.com/0987363/agent/handlers/douyu/room"
	"github.com/0987363/agent/middleware"

	"github.com/0987363/cache"

	"github.com/gin-gonic/gin"
)

var RootMux = func() *gin.Engine {
	//	gin.SetMode(gin.ReleaseMode)
	return gin.New()
}()

func init() {
	RootMux.Use(middleware.Logger())
	RootMux.Use(middleware.Recoverer())
	RootMux.Use(cache.SiteCache(nil, time.Minute*10))

	v1Mux := RootMux.Group("/v1")
	{
		douyuMux := v1Mux.Group("/douyu")
		{
			categoryMux := douyuMux.Group("/category")
			{
				categoryMux.GET("/", dcate.List)
				categoryMux.GET("/short/:short_name/sub", dcate.ListSub)
			}

			roomMux := douyuMux.Group("/room")
			{
				roomMux.GET("/", droom.List)
				roomMux.GET("/category/:id", droom.ListCategory)
				roomMux.GET("/tag/:id", droom.ListSub)

				roomMux.GET("/id/:id", droom.Get)
			}
		}
	}

	v2Mux := RootMux.Group("/v2")
	{
		douyuMux := v2Mux.Group("/douyu")
		{
			roomMux := douyuMux.Group("/room")
			{
				roomMux.GET("/id/:id", droom.GetV2)
			}
		}
	}

	v3Mux := RootMux.Group("/v3")
	{
		douyuMux := v3Mux.Group("/douyu")
		{
			roomMux := douyuMux.Group("/room")
			{
				roomMux.GET("/id/:id", droom.GetV3)
			}
		}
	}

}
