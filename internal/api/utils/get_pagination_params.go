package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPaginationParams(c *gin.Context) (int, int) {
	limit := 10 // Default limit
	page := 1   // Default page

	if l := c.Query("limit"); l != "" {
		if parsedLimit, err := strconv.Atoi(l); err == nil {
			limit = parsedLimit
		}
	}

	if p := c.Query("page"); p != "" {
		if parsedPage, err := strconv.Atoi(p); err == nil {
			page = parsedPage
		}
	}

	return limit, page
}
