package storage

type Model struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
