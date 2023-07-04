package model

type Post struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type User struct {
	ID    string  `json:"id"`
	Posts []*Post `json:"posts"`
}
