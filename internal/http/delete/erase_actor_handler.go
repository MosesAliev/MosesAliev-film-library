package delete

import (
	"bytes"
	"encoding/json"
	"film-library/internal/database"
	"film-library/internal/models"
	"net/http"
)

// http-обработчик для удаления актёра из БД
func EraseActorHandler(w http.ResponseWriter, r *http.Request) {
	if w.Header().Get("role") != "admin" {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Set("Content-Type", "text/Text")
		w.Write([]byte("нет доступа"))
		return
	}

	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	actor := models.Actor{}
	json.Unmarshal(buf.Bytes(), &actor)
	result := database.DB.Db.Where("name = ?", actor.Name).Delete(&actor) // запрос в БД для удаления информации о фильме
	if result.Error != nil {
		w.Header().Set("Content-Type", "applictaion/json")
		w.Write([]byte("Фильм не найден"))
		return
	}
	w.Write([]byte("deleted"))
}
