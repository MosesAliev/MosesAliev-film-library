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
	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	newMovie := models.Movie{}
	json.Unmarshal(buf.Bytes(), &newMovie)
	database.DB.Db.Create(&newMovie) // запрос в БД на добавление информации о фильме
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(fmt.Sprintf(`{"movie":"%s"}`, newMovie.Name)))
}
