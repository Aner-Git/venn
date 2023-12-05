package model

type Neighborhood struct {
	ID                           uint    `json:"id" gorm:"primarykey"`
	Name                         string  `json:"neighborhood" gorm:"type:varchar(128)"`
	State                        string  `json:"state" gorm:"type:varchar(128)"`
	City                         string  `json:"city" gorm:"type:varchar(128)"`
	AverageAge                   uint    `json:"average_age" gorm:"index"`
	DistanceFromCityCenter       float32 `json:"distance_from_city_center" gorm:"type:Decimal(4,1);index"`
	AverageIncome                uint    `json:"average_income"`
	Latitude                     float32 `json:"latitude" gorm:"type:Decimal(9,7)"`
	Longitude                    float32 `json:"longitude" gorm:"type:Decimal(10,7)"`
	PublicTransportAvailablityID uint    `json:"-"`
	PublicTransportAvailablity   PublicTransportAvailablity
}

type NeighborhoodExternal struct {
	ID                          uint    `json:"id" gorm:"primarykey"`
	Name                        string  `json:"neighborhood" gorm:"type:varchar(128)"`
	State                       string  `json:"state" gorm:"type:varchar(128)"`
	City                        string  `json:"city" gorm:"type:varchar(128)"`
	AverageAge                  uint    `json:"average_age" gorm:"index"`
	DistanceFromCityCenter      float32 `json:"distance_from_city_center" gorm:"type:Decimal(4,1);index"`
	AverageIncome               uint    `json:"average_income"`
	Latitude                    float32 `json:"latitude" gorm:"type:Decimal(9,7)"`
	Longitude                   float32 `json:"longitude" gorm:"type:Decimal(10,7)"`
	PublicTransportAvailability string  `json:"public_transport_availability"`
}
