package storage

import (
	"encoding/json"
	"errors"

	"github.com/aanchalverma/Machine-Coding/blogging/models"

	"github.com/go-redis/redis/v8"
)

// SaveUser saves a user in Redis.
func SaveUser(user models.User) error {
	// Convert user to JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// Save the user in Redis with the key as the username
	err = rdb.Set(ctx, "user:"+user.Username, userJSON, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

// GetUser retrieves a user by username from Redis.
func GetUser(username string) (models.User, error) {
	var user models.User

	// Get the user from Redis
	userJSON, err := rdb.Get(ctx, "user:"+username).Result()
	if err == redis.Nil {
		return user, errors.New("user not found")
	} else if err != nil {
		return user, err
	}

	// Unmarshal JSON to user struct
	err = json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		return user, err
	}

	return user, nil
}
