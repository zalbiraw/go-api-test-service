package model

type Comment struct {
	ID     string `json:"id"`
	PostID string `json:"postId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

type Post struct {
	ID       string     `json:"id"`
	Comments []*Comment `json:"comments"`
}
