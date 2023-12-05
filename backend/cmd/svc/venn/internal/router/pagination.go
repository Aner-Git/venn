package router

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	defaultPageText    = "page"
	defaultSizeText    = "pageSize"
	defaultPage        = "1"
	defaultPageSize    = "10"
	defaultMinPagesize = 10
	defaultMaxPagesize = 100
)

var (
	pageNumber = errors.New("page number must be a positive integer")
	pageSize   = errors.New("page size must be a positive integer")
)

// Create a new pagination middleware with default values
func parsePaginationDefault() gin.HandlerFunc {
	return parsePagination(
		defaultPageText,
		defaultSizeText,
		defaultPage,
		defaultPageSize,
		defaultMinPagesize,
		defaultMaxPagesize,
	)
}

func parsePagination(pageText, sizeText, defaultPage, defaultPageSize string, minPageSize, maxPageSize int) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the page from the query string and convert it to an integer
		pageStr := c.DefaultQuery(pageText, defaultPage)
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, pageNumber)
			return
		}

		// Validate for positive page number
		if page <= 0 {
			c.AbortWithError(http.StatusBadRequest, pageNumber)
			return
		}

		// Extract the size from the query string and convert it to an integer
		sizeStr := c.DefaultQuery(sizeText, defaultPageSize)
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, pageSize)
			return
		}

		// Validate for min and max page size
		switch {
		case size > maxPageSize:
			size = maxPageSize
		case size <= 0:
			size = minPageSize
		}

		// Set the page and size in the gin context
		c.Set(pageText, page)
		c.Set(sizeText, size)

		c.Next()
	}
}
