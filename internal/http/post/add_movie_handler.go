package post

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"film-library/internal/database"
	"film-library/internal/models"
	"fmt"
	"log"
	"net/http"
)

// http-обработчик для добавления информации о фильме
func AddMovieHandler(w http.ResponseWriter, r *http.Request) {
	if w.Header().Get("role") != "admin" {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Set("Content-Type", "text/Text")
		w.Write([]byte("нет доступа"))
		return
	}

	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	var newMovie models.Movie
	json.Unmarshal(buf.Bytes(), &newMovie)
	fmt.Println(newMovie.Actors)
	result := database.DB.Db.Create(&newMovie) // запрос в БД на добавление информации о фильме
	if result.Error != nil {
		w.Header().Set("Content-Type", "applictaion/json")
		w.Write([]byte("Фильм уже есть в списке"))
		return
	}

	for i := range len(newMovie.Actors) {
		log.Print("здесь")
		database.DB.Db.Exec("INSERT INTO lists (movie, actor) VALUES (@movie, @actor)", sql.Named("movie", newMovie.Name), sql.Named("actor", newMovie.Actors[i]))
	}

	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(fmt.Sprintf(`{"movie":"%s"}`, newMovie.Name)))
}
