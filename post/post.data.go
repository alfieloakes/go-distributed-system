package post

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var postList []Post

var postMap = struct {
	sync.RWMutex
	m map[int]Post
}{m: make(map[int]Post)}

func init() {
	fmt.Println("loading posts...")
	pMap, err := loadPostMap()
	postMap.m = pMap
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d posts loaded...\n", len(postMap.m))
}

func loadPostMap() (map[int]Post, error) {
	fileName := "posts.json"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file [%s] does not exist", fileName)
	}

	file, _ := ioutil.ReadFile(fileName)
	postList := make([]Post, 0)
	err = json.Unmarshal([]byte(file), &postList)
	if err != nil {
		log.Fatal(err)
	}
	postMap := make(map[int]Post)
	for i := 0; i < len(postList); i++ {
		postMap[postList[i].PostID] = postList[i]
	}
	return postMap, nil
}

func getPost(postID int) *Post {
	postMap.RLock()
	defer postMap.RUnlock()

	if post, ok := postMap.m[postID]; ok {
		return &post
	}
	return nil
}

func removePost(postId int) {
	postMap.Lock()
	defer postMap.Unlock()
	delete(postMap.m, postId)
}

func getPostList() []Post {
	postMap.RLock()
	postList := make([]Post, 0, len(postMap.m))
	for _, value := range postMap.m {
		postList = append(postList, value)
	}
	postMap.Unlock()
	return postList
}
