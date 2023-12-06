package neighborhoods

import (
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
)

var FieldsGetNeighborhoods = map[string]string{
	"neighborhood":                  "name",
	"state":                         "state",
	"city":                          "city",
	"average_age":                   "average_age",
	"distance_from_city_center":     "distance_from_city_center",
	"average_income":                "average_income",
	"latitude":                      "latitude",
	"longitude":                     "longitude",
	"public_transport_availability": "PublicTransportAvailablity__name",
}

// Build a query string and params
//
//	Filters "agerange", "maxdistance"
func GetFilters(ctx *gin.Context) ([]any, error) {
	filters := ctx.GetStringMapString("filters")
	if len(filters) == 0 {
		return []any{}, nil
	}

	query := []string{}
	values := []any{}

	if val, ok := filters["maxdistance"]; ok {
		var data []int
		err := json.Unmarshal([]byte(val), &data)
		if err != nil {
			return []any{}, nil
		}
		if len(data) != 1 {
			return []any{}, nil
		}
		query = append(query, "distance_from_city_center <= ?")
		values = append(values, data[0])
	}

	if val, ok := filters["agerange"]; ok {
		var data []int
		err := json.Unmarshal([]byte(val), &data)
		if err != nil {
			return []any{}, nil
		}

		if len(data) != 2 {
			return []any{}, nil
		}

		query = append(query, "average_age BETWEEN  ? AND ?")
		values = append(values, data[0], data[1])
	}

	ret := []any{strings.Join(query, " AND ")}

	return append(ret, values...), nil
}
