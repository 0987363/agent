package middleware

import (
	"github.com/gin-gonic/gin"
	bolt "go.etcd.io/bbolt"
)

const dbKey = "Db"

var db *bolt.DB

func ConnectBolt(path string) (err error) {
	db, err = bolt.Open(path, 0666, nil)
	if err != nil {
		return err
	}
	return
}

func Bolt() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(dbKey, db)
		c.Next()
	}
}

func GetBolt(c *gin.Context) *bolt.DB {
	if db, ok := c.Get(dbKey); ok {
		return db.(*bolt.DB)
	}

	return nil
}
