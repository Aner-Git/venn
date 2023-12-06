package router

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var filtersNeighborhood = []string{"agerange", "maxdistance"}
var filterText = "filters"

var (
	filterParseError = errors.New("failed to parse filter")
)

// Create a new filter middleware
func parseFilterNeighborhood() gin.HandlerFunc {
	return parseFilter(
		filtersNeighborhood,
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
		c.Set(filterText, filterMap)

		c.Next()
	}
}
