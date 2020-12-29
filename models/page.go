package models

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Offset int
	Limit  int
}

func DecodePage(c *gin.Context) *Page {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	if limit == 0 {
		limit = 12
	}
	return &Page{
		Limit:  limit,
		Offset: offset,
	}
}
