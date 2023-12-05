package api

import (
	"github.com/Aner-Git/venn/backend/cmd/svc/venn/internal/dal/model"
	"github.com/gin-gonic/gin"
)

type NeighborhoodDB interface {
	GetNeighborhoods(offset, limit int) ([]*model.Neighborhood, error)
	CountNeighborhood(condition ...any) (int64, error)
}

type NeighborhoodAPI struct {
	DB NeighborhoodDB
}

func (n *NeighborhoodAPI) GetClients(c *gin.Context) {
	page := c.GetInt("page")
	pageSize := c.GetInt("pageSize")
	hoods, err := n.DB.GetNeighborhoods(page, pageSize)
	if success := successOrAbort(c, 500, err); !success {
		return
	}

	cHoods, err := n.DB.CountNeighborhood()
	if success := successOrAbort(c, 500, err); !success {
		return
	}

	pagination := withPaginaion(int64(page), int64(pageSize), cHoods)
	c.JSON(200, withMeta(hoods, pagination))
}
