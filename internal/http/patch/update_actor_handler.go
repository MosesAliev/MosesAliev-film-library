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
	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	updatedActor := models.Actor{}
	json.Unmarshal(buf.Bytes(), &updatedActor)
	database.DB.Db.Save(&updatedActor) // запросв в БД на изменение данных об акётре
	w.Write([]byte("updated"))
}
