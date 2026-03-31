package entity

type Provinces struct {
	ID   int    `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
}
