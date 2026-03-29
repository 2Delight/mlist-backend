package server

type Model struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Repository string `json:"repository"`
	Version    string `json:"version"`
}
