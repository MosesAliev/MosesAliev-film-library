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
	if w.Header().Get("role") != "admin" {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Set("Content-Type", "text/Text")
		w.Write([]byte("нет доступа"))
		return
	}

	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	newActor := models.Actor{}
	json.Unmarshal(buf.Bytes(), &newActor)
	result := database.DB.Db.Create(&newActor) // запрос в БД для добавления информации об акётере
	if result.Error != nil {
		w.Header().Set("Content-Type", "applictaion/json")
		w.Write([]byte("Актёр уже есть в списке"))
		return
	}

	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(fmt.Sprintf(`{"actor":"%s"}`, newActor.Name)))
}
