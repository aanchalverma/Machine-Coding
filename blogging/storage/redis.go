package storage

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"fmt"

	"github.com/aanchalverma/Machine-Coding/blogging/models"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func SavePost(post models.Post) error {
	jsonPost, err := json.Marshal(post)
	if err != nil {
		return err
	}
	// Save to redis
	return rdb.Set(ctx, fmt.Sprintf("post:%s", post.ID), jsonPost, 0).Err()
}

func GetPostByID(id string) (models.Post, error) {
	var post models.Post
	jsonPost, err := rdb.Get(ctx, fmt.Sprintf("post:%s", id)).Result()
	if err == redis.Nil {
		return post, errors.New("post not found")
	} else if err != nil {
		return post, err
	}

	if err := json.Unmarshal([]byte(jsonPost), &post); err != nil {
		return post, err
	}

	return post, nil
}

func GetPosts(author, date string) ([]models.Post, error) {
	var posts []models.Post
	keys, err := rdb.Keys(ctx, "post:*").Result()
	if err != nil {
		return nil, err
	}

	for _, key := range keys {
		var post models.Post
		jsonPost, _ := rdb.Get(ctx, key).Result()
		json.Unmarshal([]byte(jsonPost), &post)
		// TODO: Format date
		if (author == "" || post.Author == author) && (date == "" || post.CreatedAt.Format("2006-01-02") == date) {
			posts = append(posts, post)
		}
	}

	return posts, nil
}

func UpdatePost(post models.Post) error {
	return SavePost(post)
}

func DeletePost(id string) error {
	return rdb.Del(ctx, fmt.Sprintf("post:%s", id)).Err()
}

func GetPostsWithPagination(pageStr, limitStr string) ([]models.Post, error) {
	page, err := strconv.Atoi(pageStr)
	// page size is less than 1
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10 // Default to 10 posts per page
	}

	start := (page - 1) * limit // start from the post
	keys, err := rdb.Keys(ctx, "post:*").Result()
	if err != nil {
		return nil, err
	}

	var posts []models.Post
	for i := start; i < start+limit && i < len(keys); i++ {
		var post models.Post
		jsonPost, _ := rdb.Get(ctx, keys[i]).Result()
		json.Unmarshal([]byte(jsonPost), &post)
		// Get only the number of posts within limit
		posts = append(posts, post)
	}

	return posts, nil
}
