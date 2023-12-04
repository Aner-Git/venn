package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NeighborhoodDB struct {
}

type NeighborhoodAPI struct {
	DB NeighborhoodDB
}

func (n *NeighborhoodAPI) GetClients(c *gin.Context) {

	c.String(http.StatusOK, "hoods")
}
