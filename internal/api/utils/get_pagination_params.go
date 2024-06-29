package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetPaginationParams(c *gin.Context) (int, int) {
	limit := 10 // Default limit
	page := 1   // Default page

	if l := c.Query("limit"); l != "" {
		fmt.Sscan(l, &limit)
	}

	if o := c.Query("page"); o != "" {
		fmt.Sscan(o, &page)
	}

	return limit, page
}
