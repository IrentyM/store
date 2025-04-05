package domain

type Product struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Category string
	Price    float64
	Stock    int
}
