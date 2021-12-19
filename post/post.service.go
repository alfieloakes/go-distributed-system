package post

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"
// 	"strconv"
// 	"strings"
// )

// func postHandler(w http.ResponseWriter, r *http.Request) {
// 	urlPathSegments := strings.Split(r.URL.Path, "posts/")
// 	postID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}
// 	post, listItemIndex := findPostByID(postID)
// 	if post == nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}

// 	switch r.Method {
// 	case http.MethodGet:
// 		// return the single post
// 		postJSON, err := json.Marshal(post)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(postJSON)
// 	case http.MethodPut:
// 		// update the post
// 		var updatedPost Post
// 		bodyBytes, err := ioutil.ReadAll(r.Body)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		err = json.Unmarshal(bodyBytes, &updatedPost)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		if updatedPost.PostID != postID {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		post = &updatedPost
// 		postList[listItemIndex] = *post
// 		w.WriteHeader(http.StatusOK)
// 		return
// 	default:
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		return

// 	}
// }

// func postsHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		productsJson, err := json.Marshal(postList)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(productsJson)
// 	case http.MethodPost:
// 		// Add a new post
// 		var newPost Post
// 		bodyBytes, err := ioutil.ReadAll(r.Body)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		err = json.Unmarshal(bodyBytes, &newPost)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		if newPost.PostID != 0 {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		// Store the new post in the DB
// 	}
// }
