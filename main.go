// main.go

package main

import (
	"BlogPost/db"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	dbConn, err := db.Init("host=your-db-host user=your-db-user dbname=your-db-name sslmode=disable password=your-db-password")
	if err != nil {
		log.Fatal("Failed to initialize the database: ", err)
	}
	defer db.CloseDB()

	db.AutoMigrate(&BlogPost{})
	router := mux.NewRouter()

	// Create a BlogPostService instance using the interface
	blogPostService := NewBlogPostService(db)

	// Middleware for logging requests
	router.Use(loggingMiddleware)

	// Routes
	router.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		GetPosts(w, r, blogPostService)
	}).Methods("GET")
	router.HandleFunc("/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetPost(w, r, blogPostService)
	}).Methods("GET")
	router.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		CreatePost(w, r, blogPostService)
	}).Methods("POST")
	router.HandleFunc("/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		UpdatePost(w, r, blogPostService)
	}).Methods("PUT")
	router.HandleFunc("/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		DeletePost(w, r, blogPostService)
	}).Methods("DELETE")
	router.HandleFunc("/posts/bulk", func(w http.ResponseWriter, r *http.Request) {
		CreateBulkPosts(w, r, blogPostService)
	}).Methods("POST")

	// Handling CORS (Cross-Origin Resource Sharing) for the API
	headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE"})

	// HTTP server with the CORS middleware
	http.Handle("/", handlers.CORS(headersOk, originsOk, methodsOk)(router))

	// Custom error handling
	http.Handle("/", customErrorHandler(router))

	// Serve the API on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Middleware for logging HTTP requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

// Custom error handling
func customErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovering from panic: %v", r)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
