// main.go

package main

import (
	"BlogPost/db"
	"BlogPost/handler"
	"BlogPost/service"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load configuration from environment variables
	os.Setenv("key", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	dbConnectionString := os.Getenv("key")
	port := os.Getenv("PORT")
	if dbConnectionString == "" {
		print(dbConnectionString + "=\n")
		log.Fatal("DB_CONNECTION_STRING environment variable is not set")
	}
	if port == "" {
		port = "8080" // Default port
	}

	dbConn, err := db.Init(dbConnectionString)
	if err != nil {
		log.Fatal("Failed to initialize the database: ", err)
	}
	defer dbConn.Close()

	// Create a new BlogPostService using the dbConn
	blogPostService := service.NewBlogPostService(dbConn)

	// Create a new router
	router := mux.NewRouter()

	// Use request logging middleware
	router.Use(loggingMiddleware)

	// API routes for v1
	v1 := router.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		handler.GetPosts(w, r, blogPostService)
	}).Methods("GET")
	v1.HandleFunc("/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		handler.GetPost(w, r, blogPostService)
	}).Methods("GET")
	v1.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		handler.CreatePost(w, r, blogPostService)
	}).Methods("POST")
	v1.HandleFunc("/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		handler.UpdatePost(w, r, blogPostService)
	}).Methods("PUT")
	v1.HandleFunc("/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		handler.DeletePost(w, r, blogPostService)
	}).Methods("DELETE")

	// Handling CORS (Cross-Origin Resource Sharing) for the API
	headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE"})

	// HTTP server with the CORS middleware
	http.Handle("/", handlers.CORS(headersOk, originsOk, methodsOk)(router))

	// Custom error handling
	//http.Handle("/", customErrorHandler(router))

	// Serve the API on the specified port
	addr := ":" + port
	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// Middleware for request logging
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
