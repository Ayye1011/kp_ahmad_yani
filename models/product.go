package models

type Category struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"unique" json:"category"`
}

type Product struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	Name       string
	CategoryID uint
	Category   Category `gorm:"foreignKey:CategoryID"`
}
