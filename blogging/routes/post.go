package routes

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/aanchalverma/Machine-Coding/blogging/models"
	"github.com/aanchalverma/Machine-Coding/blogging/storage"
	"github.com/aanchalverma/Machine-Coding/blogging/utils"
)

func AuthMiddleware(next http.HandlerFunc, requiredRole string) http.HandlerFunc {
	// Check if the user has write request before hitting POST
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		claims, err := utils.ValidateToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if claims.Role != requiredRole && claims.Role != "admin" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Call HandleCreatePost
		next(w, r)
	}
}
func HandleGetPosts(w http.ResponseWriter, r *http.Request) {
	// TODO: Fix this function
	// Get query parameters for pagination
	page := r.URL.Query().Get("page")
	// Enable for pagination : To be tested
	limit := r.URL.Query().Get("limit")
	posts, err := storage.GetPostsWithPagination(page, limit)
	if err != nil {
		http.Error(w, "Error retrieving posts", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(posts)
}

// Handle /posts/
func PostsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		HandleGetPosts(w, r)
	case http.MethodPost:
		AuthMiddleware(HandleCreatePost, "write")(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Handle /posts/{id}
func PostHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/posts/")
	if id == "" {
		http.Error(w, "Post ID is required", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		handleGetPostByID(w, r, id)
	case http.MethodPut:
		handleUpdatePost(w, r, id)
	case http.MethodDelete:
		handleDeletePost(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	// Validate post payload
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	post.ID = utils.GenerateID()
	post.CreatedAt = utils.GetCurrentTime() // TODO: Fix the date
	post.ModifiedAt = post.CreatedAt

	if err := storage.SavePost(post); err != nil {
		http.Error(w, "Error saving post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func handleGetPostByID(w http.ResponseWriter, r *http.Request, id string) {
	post, err := storage.GetPostByID(id)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(post)
}

func handleUpdatePost(w http.ResponseWriter, r *http.Request, id string) {
	var updatedPost models.Post
	// Validate update payload
	if err := json.NewDecoder(r.Body).Decode(&updatedPost); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedPost.ID = id
	updatedPost.ModifiedAt = utils.GetCurrentTime()

	if err := storage.UpdatePost(updatedPost); err != nil {
		http.Error(w, "Error updating post", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedPost)
}

func handleDeletePost(w http.ResponseWriter, r *http.Request, id string) {
	if err := storage.DeletePost(id); err != nil {
		http.Error(w, "Error deleting post", http.StatusInternalServerError) // 500
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
