package post

type Post struct {
	PostID int    `json:"postId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
