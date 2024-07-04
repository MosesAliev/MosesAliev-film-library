package get

import (
	"encoding/json"
	"film-library/internal/database"
	"film-library/internal/models"
	"log"
	"net/http"
)

// http-обработчик для получения списка актёров
func ActorHandler(w http.ResponseWriter, r *http.Request) {
	var actors []models.Actor               // список актёров
	database.DB.Db.Find(&actors)            // запрос в БД на получение списка актёров
	jsonActors, err := json.Marshal(actors) // список актёров в виде json
	if err != nil {
		log.Println("Ошибка десериализации")
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(jsonActors))
}
