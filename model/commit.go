package model

// Commit represents a commit in the Gitlab API
type Commit struct {
	ID             string `json:"id"`
	ShortID        string `json:"short_id"`
	Title          string `json:"title"`
	Message        string `json:"message"`
	CreatedAt      string `json:"created_at"`
	AuthorName     string `json:"author_name"`
	AuthorEmail    string `json:"author_email"`
	CommitterName  string `json:"committer_name"`
	CommitterEmail string `json:"committer_email"`
}
