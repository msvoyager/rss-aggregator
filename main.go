package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/msvoyager/rss-aggregator/internal/database"
    _ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	

	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}


	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("cant connect to the database", err)

	}

	// queries, err := database.new(conn)
	// if err != nil {
	// 	log.Fatal("cant create to db connection", err)
	// }

	apiCfg := apiConfig{
		DB: database.New(conn),
	}
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"https://*", "http://*"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
  }))

  	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
  	v1Router.Get("/err", handlerError)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)
	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	router.Mount("/v1", v1Router)


	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,

	}

	log.Printf("\nserver starting on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PORT:", portString)

}