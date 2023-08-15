package dbmodel

type Category struct {
	ID        int    `gorm:"primary_key;not null"`
	CompanyID int    `gorm:"type:int(11);not null"`
	Name      string `gorm:"type:varchar(255);not null"`
}

func (c Category) TableName() string {
	return "category"
}
