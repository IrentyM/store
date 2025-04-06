package domain

type Category struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Dscription string `db:"description"`
}
