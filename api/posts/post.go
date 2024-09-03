package posts

import (
	"errors"
	"fmt"

	"github.com/jeremitraverse/explo/db"
	"github.com/jeremitraverse/explo/types"
)

func CreatePost(post types.Post) {
	con, err := db.CreateConnection()

	if post.Title == "" || post.Content == "" {
		return
	}

	if err != nil {
		fmt.Println("Post not create: ", err)
	}

	con.MustExec("INSERT INTO posts (title, content) VALUES ($1, $2)", post.Title, post.Content)

	con.Close()
}

func GetPosts() []types.Post {
	posts := []types.Post{}

	con, err := db.CreateConnection()

	if err != nil {
		fmt.Println("Post not create")
	}

	err = con.Select(&posts, "SELECT * FROM posts")

	if err != nil {
		fmt.Println("Post not create: ", err)
	}

	con.Close()

	return posts
}

func GetPost(id string) (types.Post, error) {
	post := types.Post{}

	con, err := db.CreateConnection()

	if err != nil {
		fmt.Println("error while creating a connection to db: ", err)
		return post, err
	}

	err = con.Get(&post, "SELECT * FROM posts WHERE post_key = $1", id)

	if err != nil {
		fmt.Println("error while fetching post from db: ", err)
		return post, err
	}

	return post, nil
}

func DeletePost(id string) (string, error) {
	con, err := db.CreateConnection()

	if err != nil {
		fmt.Println("error while creating a connection to db: ", err)
		return "", err
	}

	_, err = con.Exec("DELETE FROM posts WHERE post_key = $1", id)

	if err != nil {
		fmt.Println("error while fetching post from db: ", err)
		return "", err
	}

	return id, nil
}

func UpdatePost(post types.Post) (string, error) {
	con, err := db.CreateConnection()

	if err != nil {
		errorMessage := fmt.Sprintf("error while creating a new connection to db: %s", err.Error())
		return "", errors.New(errorMessage)
	}

	if post.PostKey == "" || post.Content == "" || post.Title == "" {
		return "", errors.New("post title and post content can't be empty")
	}

	_, err = con.Exec("UPDATE posts SET content = $1, title = $1  WHERE post_key = $1", post.PostKey)

	if err != nil {
		errorMessage := fmt.Sprintf("error while updating post: %s", err.Error())
		return "", errors.New(errorMessage)
	}

	con.Close()

	return post.PostKey, nil
}
