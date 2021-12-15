package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Post struct {
	PostID int    `json:"postId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

var postList []Post

func init() {
	postJSON := `[
		{
			"postId": 1,
			"title": "Hello, World!",
			"body": "It's nice to meet you!"
		},
		{
			"postId": 2,
			"title": "Post Number 2",
			"body": "It's nice to see you again!"
		}
	]`

	err := json.Unmarshal([]byte(postJSON), &postList)
	if err != nil {
		log.Fatal(err)
	}
}

func findPostByID(postID int) (*Post, int) {
	for i, post := range postList {
		if post.PostID == postID {
			return &post, i
		}
	}

	return nil, 0
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "posts/")
	postID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	post, listItemIndex := findPostByID(postID)
	if post == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// return the single post
		postJSON, err := json.Marshal(post)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(postJSON)
	case http.MethodPut:
		// update the post
		var updatedPost Post
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &updatedPost)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if updatedPost.PostID != postID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		post = &updatedPost
		postList[listItemIndex] = *post
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return

	}
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productsJson, err := json.Marshal(postList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productsJson)
	case http.MethodPost:
		// Add a new post
		var newPost Post
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &newPost)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if newPost.PostID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// Store the new post in the DB
	}
}

func main() {
	http.HandleFunc("/posts", postsHandler)
	http.HandleFunc("/posts/", postHandler)
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
