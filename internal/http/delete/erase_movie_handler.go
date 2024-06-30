package delete

import (
	"bytes"
	"encoding/json"
	"film-library/internal/database"
	"film-library/internal/models"
	"net/http"
)

func EraseMovieHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	deletedMovie := models.Movie{}
	json.Unmarshal(buf.Bytes(), &deletedMovie)
	database.DB.Db.Delete(&deletedMovie)
	w.Write([]byte("deleted"))
}
