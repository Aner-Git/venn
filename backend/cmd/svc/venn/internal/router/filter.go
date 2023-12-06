package router

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var filters = []string{"agerange", "maxdistance"}
var filterText = "filters"

var (
	filterParseError = errors.New("failed to parse filter")
)

// Create a new filter middleware
func parseFilterDefault() gin.HandlerFunc {
	return parseFilter(
		filters,
	)
}

func parseFilter(filters []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		filterMap := make(map[string]string)
		// Extract the filters query
		for _, filterName := range filters {
			fStr := c.Query(filterName)
			if fStr != "" {
				filterMap[filterName] = fStr
			}
		}

		// Set filters the gin context
		if len(filters) != 0 {
			c.Set(filterText, filterMap)
		}

		c.Next()
	}
}
