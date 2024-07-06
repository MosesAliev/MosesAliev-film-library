package patch

import (
	"bytes"
	"encoding/json"
	"film-library/internal/database"
	"film-library/internal/models"
	"net/http"
)

// http-обработчик для обновления информации о фильмах
func UpdateMovieHandler(w http.ResponseWriter, r *http.Request) {
	if w.Header().Get("role") != "admin" {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Set("Content-Type", "text/Text")
		w.Write([]byte("нет доступа"))
		return
	}

	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	updatedMovie := models.Movie{}
	json.Unmarshal(buf.Bytes(), &updatedMovie)
	result := database.DB.Db.Save(&updatedMovie) // запрос в БД на изменение данных о фильме
	if result.Error != nil {
		w.Header().Set("Content-Type", "applictaion/json")
		w.Write([]byte("Фильм не найден"))
		return
	}

	w.Write([]byte("updated"))
}
