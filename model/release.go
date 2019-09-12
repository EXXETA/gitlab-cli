package model

// Release represents a release in the Gitlab API
type Release struct {
	Name        string `json:"name"`
	TagName     string `json:"tag_name"`
	Description string `json:"description"`
	Reference   string `json:"ref"`
	Author      User   `json:"author"`
	Commit      Commit `json:"commit"`
	Assets      asset  `json:"assets"`
	ReleasedAt  string `json:"released_at"`
	CreatedAt   string `json:"created_at"`
}

type asset struct {
	Count   int      `json:"count"`
	Sources []source `json:"sources"`
	Links   []link   `json:"links"`
}

type link struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type source struct {
	Format string `json:"format"`
	URL    string `json:"url"`
}
