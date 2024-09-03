package posts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jeremitraverse/explo/types"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)

	var post types.Post
	err := decoder.Decode(&post)

	if err != nil {
		fmt.Println("error while decoding post from request: ", err)
		return
	}

	CreatePost(post)
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	posts := GetPosts()

	jsonPosts, err := json.Marshal(posts)

	if err != nil {
		fmt.Println("error marshaling posts: ", err)
	}

	w.Write(jsonPosts)
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("id")

	if key == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	post, err := GetPost(key)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	jsonPost, err := json.Marshal(post)

	if err != nil {
		fmt.Println("error marshaling posts: ", err)
		return
	}

	w.Write(jsonPost)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	if id == "" {
		createBadRequest(w, "post id must be provided")
		return
	}

	_, err := DeletePost(id)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// TODO: Create json response with success
}

func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	decoder := json.NewDecoder(r.Body)

	var post types.Post
	err := decoder.Decode(&post)

	if err != nil {
		createBadRequest(w, err.Error())
	}

	if post.PostKey == "" || post.Content == "" || post.Title == "" {
		http.Error(w, "post key, post content and post title can't be empty", http.StatusBadRequest)
		return
	}

	_, err = UpdatePost(post)

	if err != nil {
		createBadRequest(w, err.Error())
		return
	}

}

func createBadRequest(w http.ResponseWriter, errorMessage string) {
	formattedMessage := fmt.Sprintf("Bad request:  %s", errorMessage)
	http.Error(w, formattedMessage, http.StatusBadRequest)
}
