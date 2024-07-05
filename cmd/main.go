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
	database.ConnectDB() // подключение к БД
	http.Handle("POST /addactor", http.HandlerFunc(post.Auth(post.AddActorHandler)))
	http.Handle("POST /addmovie", http.HandlerFunc(post.Auth(post.AddMovieHandler)))
	http.Handle("GET /actor", http.HandlerFunc(get.ActorHandler))
	http.Handle("GET /movie", http.HandlerFunc(post.Auth(get.MovieHandler)))
	http.Handle("PATCH /updateactor", http.HandlerFunc(post.Auth(patch.UpdateActorHandler)))
	http.Handle("PATCH /updatemovie", http.HandlerFunc(post.Auth(patch.UpdateMovieHandler)))
	http.Handle("DELETE /eraseactor", http.HandlerFunc(post.Auth(delete.EraseActorHandler)))
	http.Handle("DELETE /erasemovie", http.HandlerFunc(post.Auth(delete.EraseMovieHandler)))
	http.ListenAndServe(":8080", nil) // сервер
}
