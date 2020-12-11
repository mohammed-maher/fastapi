package requests

import (
	"errors"
	"time"
)

type TripRequest struct {
	CarID         uint     `json:"car_id"`
	FromCityID    uint     `json:"from_city_id"`
	ToCityID      uint     `json:"to_city_id"`
	Passengers    uint        `json:"passengers"`
	IsTwoWay      bool       `json:"two_way"`
	DepartureDate time.Time  `json:"departure_date"`
	ReturnDate    *time.Time `json:"return_date"`
}

func (r *TripRequest) Validate(isUpdate bool) error {
	params:=map[string]interface{}{
		"car_id":r.CarID,
		"from_city_id":r.FromCityID,
		"to_city_id":r.ToCityID,
		"passengers":r.Passengers,
		"two_way":r.IsTwoWay,
		"departure_date":r.DepartureDate,
		"return_date":r.ReturnDate,
	}
	for k,v:=range params{
		if isUpdate&&v==nil{
			continue
		}
		if v==nil{
			return errors.New("invalid_"+k)
		}
	}
	return nil
}
