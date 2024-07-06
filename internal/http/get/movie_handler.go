package get

import (
	"encoding/json"
	"film-library/internal/database"
	"film-library/internal/models"
	"log"
	"net/http"
)

// http-обработчик запроса на получение списка фильмов
func MovieHandler(w http.ResponseWriter, r *http.Request) {
	listOrder, orderOk := r.URL.Query()["order"] // порядок, по котрому фильмы сортируются
	search, searchOk := r.URL.Query()["search"]  // поиск по запросу, если запроса нет, то searchOk равен false, следовательно искать ничего не нужно
	if orderOk && listOrder[0] != "rating" {     // обработка случаев, когда сортировка не по рейтингу
		if listOrder[0] == "name" {
			var movies []models.Movie
			if searchOk { // если есть запрос на поиск
				database.DB.Db.Order("name").Where("name LIKE ?", `%`+search[0]+`%`).Find(&movies) // запрос в БД с поиском и сортировкой по имени
			} else {
				database.DB.Db.Order("name").Find(&movies) // запрос в БД с сортировкой по имени
			}

			jsonMovies, err := json.Marshal(movies) // список фильмов в виде json
			if err != nil {
				log.Println("Ошибка десериализации")
				w.WriteHeader(http.StatusConflict)
				return
			}

			w.Header().Set("Content-Type", "applictaion/json")
			w.Write([]byte(jsonMovies))
		} else if listOrder[0] == "date" {
			var movies []models.Movie
			if searchOk {
				// запрос в БД с поиском и сортировкой по дате выпуска
				database.DB.Db.Order("realize_date desc").Where("name LIKE ?", `%`+search[0]+`%`).Find(&movies)
			} else {
				// запрос в БД с сортировкой по дате выпуска
				database.DB.Db.Order("realize_date desc").Find(&movies)
			}

			jsonMovies, err := json.Marshal(movies)
			if err != nil {
				log.Println("Ошибка десериализации")
				w.WriteHeader(http.StatusConflict)
				return
			}

			w.Header().Set("Content-Type", "applictaion/json")
			w.Write([]byte(jsonMovies))
		} else { // если в параметрах сортировки указано что-то лишнее, то выводится ошибка 400
			w.Header().Set("Content-Type", "applictaion/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("что то пошло не так"))
		}

	} else { // если не указан параметр сортировки, то сортировка идёт по рейтингу
		var movies []models.Movie
		if searchOk {
			// запрос в БД с поиском и сортировкой по рейтингу
			database.DB.Db.Order("rating").Where("name LIKE ?", `%`+search[0]+`%`).Find(&movies)
		} else {
			// запрос в БД с сортировкой по рейтингу
			database.DB.Db.Order("rating").Find(&movies)
		}

		jsonMovies, err := json.Marshal(movies)
		if err != nil {
			log.Println("Ошибка десериализации")
			w.WriteHeader(http.StatusConflict)
			return
		}

		w.Header().Set("Content-Type", "applictaion/json")
		w.Write([]byte(jsonMovies))
	}

}
