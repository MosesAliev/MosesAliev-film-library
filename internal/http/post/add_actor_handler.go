package post

import (
	"bytes"
	"encoding/json"
	"film-library/internal/database"
	"film-library/internal/models"
	"fmt"
	"net/http"
)

func AddActorHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	newActor := models.Actor{}
	json.Unmarshal(buf.Bytes(), &newActor)
	database.DB.Db.Create(&newActor)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(fmt.Sprintf(`{"actor":"%s"}`, newActor.Name)))
}
