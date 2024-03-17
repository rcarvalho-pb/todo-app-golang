package cookies

import (
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/rcarvalho-pb/todo-app-golang/internal/config"
)

var s *securecookie.SecureCookie

func Config() {
	s = securecookie.New(config.EnvConfigs.HashKey, config.EnvConfigs.BlockKey)
}

func Save(w http.ResponseWriter, id, role, token string) error {
	data := map[string]string{
		"id":    id,
		"role":  role,
		"token": token,
	}

	encodedData, err := s.Encode(config.EnvConfigs.CookieName, data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     config.EnvConfigs.CookieName,
		Value:    encodedData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

func Read(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie(config.EnvConfigs.CookieName)
	if err != nil {
		return nil, err
	}

	var values map[string]string

	if err = s.Decode(config.EnvConfigs.CookieName, cookie.Name, &values); err != nil {
		return nil, err
	}

	return values, nil
}

func Delete(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     config.EnvConfigs.CookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})
}
