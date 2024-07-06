package post

import (
	"film-library/internal/database"
	"film-library/internal/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.Split(r.Header.Get("Authorization"), " ")[1]
		fmt.Println(token)
		secretKey := []byte("auth")
		jwtToken, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		res, ok := jwtToken.Claims.(jwt.MapClaims)
		// обязательно используем второе возвращаемое значение ok и проверяем его, потому что
		// если Сlaims вдруг оказжется другого типа, мы получим панику
		if !ok {
			fmt.Printf("failed to typecast to jwt.MapCalims")
			return
		}

		loginRaw := res["login"]
		rolesRaw := res["role"]
		// loginRaw — интерфейс, так как тип значения в jwt.Claims — интерфейс. Чтобы получить строку, нужно
		// снова сделать приведение типа к строке.
		login, ok := loginRaw.(string)
		if !ok {
			fmt.Printf("failed to typecast to string login")
			return
		}
		// обратите внимание, что при создании мы указывали тип []string, однако тут приводим к []inteface{}
		// так происходит, потому что json не строго типизированный, из-за чего при парсинге нельзя точно
		// определить тип слайса.
		role, ok := rolesRaw.(string)
		if !ok {
			fmt.Printf("failed to typecast to string role")
			return
		}

		user := models.User{}

		database.DB.Db.First(&user, "login = ? AND role = ?", login, role)
		fmt.Println(user.Login)
		fmt.Println(user.Role)
		if len(user.Login) == 0 && len(user.Role) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "applictaion/json")
			w.Write([]byte("не авторизован"))
			return
		}

		w.Header().Add("role", user.Role)
		next(w, r)
	})

}
