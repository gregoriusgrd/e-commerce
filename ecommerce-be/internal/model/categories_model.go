package model

type CategoryResponse struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Children []CategoryResponse `json:"children"`
}