package delete

import (
	"bytes"
	"encoding/json"
	"film-library/internal/database"
	"film-library/internal/models"
	"net/http"
)

func EraseActorHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	actor := models.Actor{}
	json.Unmarshal(buf.Bytes(), &actor)
	database.DB.Db.Delete(&actor)
	w.Write([]byte("deleted"))
}
