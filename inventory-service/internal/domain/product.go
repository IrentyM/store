package domain

import (
	"fmt"
	"time"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	CategoryID  int
	Stock       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (p *Product) Validate() error {
	switch {
	case p.Name == "":
		return fmt.Errorf("product name cannot be empty")
	case p.Price < 0:
		return fmt.Errorf("product price cannot be negative")
	case p.CategoryID <= 0:
		return fmt.Errorf("invalid category ID")
	case p.Stock < 0:
		return fmt.Errorf("product stock cannot be negative")
	default:
		return nil
	}
}
