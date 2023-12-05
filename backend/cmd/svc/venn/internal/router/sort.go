package router

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Aner-Git/venn/backend/cmd/svc/venn/internal/common"
	"github.com/gin-gonic/gin"
)

const (
	defaultSortText = "sort"
)

var (
	sortDidNotParse = errors.New("failed to parse sort")
)

// Create a new sort parse middleware
func parseSortDefault() gin.HandlerFunc {
	return parseSort(
		defaultSortText,
	)
}

func parseSort(sortText string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the sort query
		sortStr := c.DefaultQuery(sortText, "[]")

		var sort []*common.SortBy
		err := json.Unmarshal([]byte(sortStr), &sort)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, sortDidNotParse)
			return
		}

		c.Set(sortText, nil)
		// Set the sort in the gin context
		if len(sort) != 0 {
			c.Set(sortText, sort[0])
		}

		c.Next()
	}
}
