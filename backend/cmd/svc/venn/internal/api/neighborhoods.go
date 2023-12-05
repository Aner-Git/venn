package api

import (
	"github.com/Aner-Git/venn/backend/cmd/svc/venn/internal/api/neighborhoods"
	"github.com/Aner-Git/venn/backend/cmd/svc/venn/internal/dal/model"
	"github.com/gin-gonic/gin"
)

type NeighborhoodDB interface {
	GetNeighborhoods(orderBy string, offset, limit int) ([]*model.Neighborhood, error)
	CountNeighborhood(condition ...any) (int64, error)
}

type NeighborhoodAPI struct {
	DB NeighborhoodDB
}

func (n *NeighborhoodAPI) GetNeighborhoods(c *gin.Context) {
	page := c.GetInt("page")
	pageSize := c.GetInt("pageSize")

	sortBy := getSort(c)
	orderClause := getOrderClause(sortBy, neighborhoods.FieldsGetNeighborhoods, "name")
	hoods, err := n.DB.GetNeighborhoods(orderClause, page, pageSize)
	if success := successOrAbort(c, 500, err); !success {
		return
	}

	cntHoods, err := n.DB.CountNeighborhood()
	if success := successOrAbort(c, 500, err); !success {
		return
	}

	pagination := withPaginaion(int64(page), int64(pageSize), cntHoods)
	c.JSON(200, withMeta(neighborhoods.ToExternal(hoods), pagination))
}
