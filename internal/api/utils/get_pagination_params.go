package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetPaginationParams(c *gin.Context) (int, int) {
	limit := 10 // Default limit
	offset := 0 // Default offset

	if l := c.Query("limit"); l != "" {
		fmt.Sscan(l, &limit)
	}

	if o := c.Query("offset"); o != "" {
		fmt.Sscan(o, &offset)
	}

	return limit, offset
}
