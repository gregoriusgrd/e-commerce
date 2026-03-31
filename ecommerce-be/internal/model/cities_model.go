package model

type CitiesResponse struct{
	ID int `json:"id"`
	Name string `json:"name"`
}

type GetCityRequest struct{
	ProvinceId int `query:"province_id" validate:"required"`
}