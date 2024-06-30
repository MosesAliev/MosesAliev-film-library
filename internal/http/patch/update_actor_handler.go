package patch

import (
	"bytes"
	"encoding/json"
	"film-library/internal/database"
	"film-library/internal/models"
	"net/http"
)

func UpdateActorHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	updatedActor := models.Actor{}
	json.Unmarshal(buf.Bytes(), &updatedActor)
	database.DB.Db.Save(&updatedActor)
	w.Write([]byte("updated"))
}
