package patch

import (
	"bytes"
	"encoding/json"
	"film-library/internal/database"
	"film-library/internal/models"
	"net/http"
)

func UpdateMovieHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	updatedMovie := models.Movie{}
	json.Unmarshal(buf.Bytes(), &updatedMovie)
	database.DB.Db.Save(&updatedMovie)
	w.Write([]byte("updated"))
}
