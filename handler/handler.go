// handler/handler.go

package handler

import (
	"BlogPost/model"
	"BlogPost/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetPosts(w http.ResponseWriter, r *http.Request, blogPostService service.BlogPostService) {
	posts, err := blogPostService.GetAllPosts()
	if err != nil {
		http.Error(w, "Failed to retrieve blog posts", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, posts)
}

func GetPost(w http.ResponseWriter, r *http.Request, blogPostService service.BlogPostService) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := blogPostService.GetPostByID(uint(id))
	if err != nil {
		if err == service.ErrPostNotFound {
			http.Error(w, "Post not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve the post", http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, http.StatusOK, post)
}

func CreatePost(w http.ResponseWriter, r *http.Request, blogPostService service.BlogPostService) {
	var post model.BlogPost
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := blogPostService.CreatePost(&post); err != nil {
		http.Error(w, "Failed to create the post", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request, blogPostService service.BlogPostService) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var updatedPost model.BlogPost
	if err := json.NewDecoder(r.Body).Decode(&updatedPost); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := blogPostService.UpdatePost(uint(id), &updatedPost); err != nil {
		if err == service.ErrPostNotFound {
			http.Error(w, "Post not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to update the post", http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, http.StatusOK, updatedPost)
}

func DeletePost(w http.ResponseWriter, r *http.Request, blogPostService service.BlogPostService) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	if err := blogPostService.DeletePost(uint(id)); err != nil {
		if err == service.ErrPostNotFound {
			http.Error(w, "Post not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to delete the post", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// respondWithJSON is a helper function to send JSON responses.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Failed to marshal response data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
