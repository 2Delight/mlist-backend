package storage

type Model struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Repository string `db:"repository"`
	Version    string `db:"version"`
}
