package entity

type Cities struct{
	ID int `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
	ProvinceId int `gorm:"column:province_id"`
}