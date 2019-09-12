package model

// Project represents a project in the Gitlab API
type Project struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	NameWithNameSpace string `json:"name_with_namespace"`
	CreatedAt         string `json:"created_at"`
	WebURL            string `json:"web_url"`
}
