package model

type PublicTransportAvailablity struct {
	ID   uint   `json:"id" gorm:"primarykey"`
	Name string `json:"public_transport_availability" gorm:"type:varchar(32);unique"`
}
