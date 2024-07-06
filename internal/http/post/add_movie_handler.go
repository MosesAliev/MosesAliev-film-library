package post

import (
	"bytes"
	"encoding/json"
	"film-library/internal/database"
	"film-library/internal/models"
	"fmt"
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
	newMovie := models.Movie{}
	json.Unmarshal(buf.Bytes(), &newMovie)
	result := database.DB.Db.Create(&newMovie) // запрос в БД на добавление информации о фильме
	if result.Error != nil {
		w.Header().Set("Content-Type", "applictaion/json")
		w.Write([]byte("Фильм уже есть в списке"))
		return
	}

	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(fmt.Sprintf(`{"movie":"%s"}`, newMovie.Name)))
}
