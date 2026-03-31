package entity

type Categories struct {
	ID       int          `gorm:"column:id;primaryKey"`
	Name     string       `gorm:"column:name"`
	ParentId *int         `gorm:"column:parent_id"`
	Parent   *Categories  `gorm:"foreignKey:ParentId;references:id"`
	Children []Categories `gorm:"foreignKey:ParentId;references:id"`
}
