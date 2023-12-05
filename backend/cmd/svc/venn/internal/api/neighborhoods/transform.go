package neighborhoods

import "github.com/Aner-Git/venn/backend/cmd/svc/venn/internal/dal/model"

func ToExternal(hoods []*model.Neighborhood) []*model.NeighborhoodExternal {
	ext := make([]*model.NeighborhoodExternal, len(hoods))

	for index, p := range hoods {
		h := &model.NeighborhoodExternal{
			ID:                          p.ID,
			Name:                        p.Name,
			State:                       p.State,
			City:                        p.City,
			AverageAge:                  p.AverageAge,
			DistanceFromCityCenter:      p.DistanceFromCityCenter,
			AverageIncome:               p.AverageIncome,
			Latitude:                    p.Latitude,
			Longitude:                   p.Latitude,
			PublicTransportAvailability: p.PublicTransportAvailablity.Name,
		}
		ext[index] = h
	}
	return ext
}
