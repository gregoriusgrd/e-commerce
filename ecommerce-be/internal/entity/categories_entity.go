package entity

type Categories struct {
	ID       int          `gorm:"column:id;primaryKey"`
	Name     string       `gorm:"column:name"`
	ParentId *int         `gorm:"column:parent_id"`
	Parent   *Categories  `gorm:"foreignKey:parent_id;references:id"`
	Children []Categories `gorm:"foreignKey:parent_id;references:id"`
}
