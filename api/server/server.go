package server

import (
	"fmt"
	"io"
	"net/http"

	"github.com/jeremitraverse/explo/posts"
)

func Start() {
	fmt.Println("Starting server on port 8080")

	http.HandleFunc("/", getRoot)

	http.HandleFunc("/create-post", posts.CreatePostHandler)
	http.HandleFunc("/posts", posts.GetPostsHandler)
	http.HandleFunc("/post", posts.GetPostHandler)
	http.HandleFunc("/update-post", posts.UpdatePostHandler)
	http.HandleFunc("/delete-post", posts.DeletePostHandler)

	http.ListenAndServe(":8080", nil)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Yo this is rooot")
}
