package model

type Author struct {
	Name      string `json:"name"`
	DateBirth string `json:"date_birth"`
	Bio       string `json:"bio"`
}

type CreateBlogPost struct {
	AuthorName string `json:"authorName"`
	Title      string `json:"title"`
	Text       string `json:"text"`
	Photos     []struct {
		Caption string `json:"caption"`
		Photo   string `json:"photo"`
	} `json:"photos"`
}

type BlogPost struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Text   string `json:"text"`
	Author struct {
		Name string `json:"name"`
		Bio  string `json:"bio"`
	} `json:"author"`
	Photos []struct {
		Caption string `json:"caption"`
		Photo   string `json:"photo"`
	} `json:"photos"`
	CreatedAt string `json:"created_at"`
}
