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
	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	actor := models.Actor{}
	json.Unmarshal(buf.Bytes(), &actor)
	database.DB.Db.Delete(&actor) // запрос в БД на удаление информации об актёре
	w.Write([]byte("deleted"))
}
