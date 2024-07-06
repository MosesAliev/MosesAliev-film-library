package patch

import (
	"bytes"
	"encoding/json"
	"film-library/internal/database"
	"film-library/internal/models"
	"net/http"
)

// http-обработчик для изменения информации об актёрах
func UpdateActorHandler(w http.ResponseWriter, r *http.Request) {
	if w.Header().Get("role") != "admin" {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Set("Content-Type", "text/Text")
		w.Write([]byte("нет доступа"))
		return
	}

	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	updatedActor := models.Actor{}
	json.Unmarshal(buf.Bytes(), &updatedActor)
	result := database.DB.Db.Save(&updatedActor) // запросв в БД на изменение данных об акётре
	if result.Error != nil {
		w.Header().Set("Content-Type", "applictaion/json")
		w.Write([]byte("Актёр не найден"))
		return
	}

	w.Write([]byte("updated"))
}
