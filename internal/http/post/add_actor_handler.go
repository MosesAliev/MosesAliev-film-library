package post

import (
	"bytes"
	"encoding/json"
	"film-library/internal/database"
	"film-library/internal/models"
	"fmt"
	"net/http"
)

// http-обработчик для добавления информации об актёре
func AddActorHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	newActor := models.Actor{}
	json.Unmarshal(buf.Bytes(), &newActor)
	database.DB.Db.Create(&newActor) // запрос в БД для добавления информации об акётере
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(fmt.Sprintf(`{"actor":"%s"}`, newActor.Name)))
}
