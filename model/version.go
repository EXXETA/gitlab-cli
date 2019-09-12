package model

// Version represents a version in the Gitlab API
type Version struct {
	Version  string `json:"version"`
	Revision string `json:"revision"`
}
