package domain

import "fmt"

type Category struct {
	ID         int32  `db:"id"`
	Name       string `db:"name"`
	Dscription string `db:"description"`
}

func (c *Category) Validate() error {
	switch {
	case c.Name == "":
		return fmt.Errorf("category name cannot be empty")
	case len(c.Name) > 255:
		return fmt.Errorf("category name cannot exceed 255 characters")
	default:
		return nil
	}
}
