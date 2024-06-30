package main

import (
	"film-library/internal/database"
	"film-library/internal/http/delete"
	"film-library/internal/http/get"
	"film-library/internal/http/patch"
	"film-library/internal/http/post"
	"net/http"
)

func main() {
	database.ConnectDB()
	http.Handle("POST /addactor", http.HandlerFunc(post.AddActorHandler))
	http.Handle("POST /addmovie", http.HandlerFunc(post.AddMovieHandler))
	http.Handle("GET /actor", http.HandlerFunc(get.ActorHandler))
	http.Handle("GET /movie", http.HandlerFunc(get.MovieHandler))
	http.Handle("PATCH /updateactor", http.HandlerFunc(patch.UpdateActorHandler))
	http.Handle("PATCH /updatemovie", http.HandlerFunc(patch.UpdateMovieHandler))
	http.Handle("DELETE /eraseactor", http.HandlerFunc(delete.EraseActorHandler))
	http.Handle("DELETE /erasemovie", http.HandlerFunc(delete.EraseMovieHandler))
	http.ListenAndServe(":8080", nil)
}
