package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// type Post struct {
// 	PostID int    `json:"postId"`
// 	Title  string `json:"title"`
// 	Body   string `json:"body"`
// }

// var postList []Post

// func init() {
// 	postJSON := `[
// 		{
// 			"postId": 1,
// 			"title": "Hello, World!",
// 			"body": "It's nice to meet you!"
// 		},
// 		{
// 			"postId": 2,
// 			"title": "Post Number 2",
// 			"body": "It's nice to see you again!"
// 		}
// 	]`

// 	err := json.Unmarshal([]byte(postJSON), &postList)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func middlewareHandler(handler http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("Before handler, middleware start")
// 		handler.ServeHTTP(w, r)
// 		fmt.Println("After handler, middleware finsihed")
// 	})
// }

// func main() {
// 	postListHandler := http.HandlerFunc(po)
// 	postItemHandler := http.HandlerFunc(postHandler)

// 	http.Handle("/posts", middlewareHandler(postListHandler))
// 	http.Handle("/posts/", middlewareHandler(postItemHandler))
// 	err := http.ListenAndServe(":3000", nil)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
