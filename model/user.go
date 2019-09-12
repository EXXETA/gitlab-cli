package model

// User represents a user in the Gitlab API
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	State    string `json:"state"`
}
