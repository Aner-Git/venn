package api

import (
	"github.com/Aner-Git/venn/backend/cmd/svc/venn/internal/api/neighborhoods"
	"github.com/Aner-Git/venn/backend/cmd/svc/venn/internal/dal/model"
	"github.com/gin-gonic/gin"
)

type NeighborhoodDB interface {
	GetNeighborhoods(orderBy string, offset, limit int, condition ...any) ([]*model.Neighborhood, error)
	CountNeighborhood(condition ...any) (int64, error)
}

type NeighborhoodAPI struct {
	DB NeighborhoodDB
}

/*
*
* /neighborhoods
*
* Returns a list of neighborhoods: by default return a max of the last pageSize records(by neighborhood name desc)
*
* requrest params:
*  - pagination, interger types: page, pageSize. first page starts at 1(page=1).  e.g page=1&pageSize=33
*  - sort, json string enconded array of an object {id:"field_name","desc":true|false}. eg.  sort=[{"id":"average_age","desc":true}]
*  - maxdistance, array of integer type. e.g. maxdisnace=[25]
*  - agerange, array of 2 integers type. e.g agerange=[12,34]
*
* return a json object with data field and meta field
* - data: array of neighborhoods objects. e.g. data:[{<neighborhoods fields>}]
* - meta: an object with pagination data. e.g meta:{"page":1,"page_size":20,"total":1000}
*
*
* error:
* the endpoint returns 400 or 500 http code on errors
* Text in the body describes more info of the error
 */
func (n *NeighborhoodAPI) GetNeighborhoods(c *gin.Context) {

	page := c.GetInt("page")
	pageSize := c.GetInt("pageSize")

	sortBy := getSort(c)
	orderClause := getOrderClause(sortBy, neighborhoods.FieldsGetNeighborhoods, "name")
	filters, err := neighborhoods.GetFilters(c)
	if success := successOrAbort(c, 400, err); !success {
		return
	}
	hoods, err := n.DB.GetNeighborhoods(orderClause, page, pageSize, filters...)
	if success := successOrAbort(c, 500, err); !success {
		return
	}

	cntHoods, err := n.DB.CountNeighborhood(filters...)
	if success := successOrAbort(c, 500, err); !success {
		return
	}

	pagination := withPaginaion(int64(page), int64(pageSize), cntHoods)
	c.JSON(200, withMeta(neighborhoods.ToExternal(hoods), pagination))
}
