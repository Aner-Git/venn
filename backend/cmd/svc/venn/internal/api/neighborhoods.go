package api

import (
	"github.com/Aner-Git/venn/backend/cmd/svc/venn/internal/dal/model"
	"github.com/gin-gonic/gin"
)

type NeighborhoodDB interface {
	GetNeighborhoods() ([]*model.Neighborhood, error)
}

type NeighborhoodAPI struct {
	DB NeighborhoodDB
}

func (n *NeighborhoodAPI) GetClients(c *gin.Context) {

	hoods, err := n.DB.GetNeighborhoods()
	if success := successOrAbort(c, 500, err); !success {
		return
	}
	c.JSON(200, hoods)
}
