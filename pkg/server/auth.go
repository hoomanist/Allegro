package server

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/hoomanist/allegro-server/pkg/auth"
	"github.com/hoomanist/allegro-server/pkg/database"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *server) Login(w http.ResponseWriter, r *http.Request) {
	var m loginRequest
	json.NewDecoder(r.Body).Decode(&m)
	hashed_password := Hash([]byte(m.Password))
	pswd, err := database.GetUser(s.SqlCfg, m.Username)
	switch err {
	case sql.ErrNoRows:
		http.Error(w, err.Error(), http.StatusNoContent)
	case nil:
		if pswd == hashed_password {
			// proceed with login
			t := auth.NewToken(m.Username, s.Key)
			json.NewEncoder(w).Encode(map[string]string{
				"token": t,
			})
		} else {
			http.Error(w, errors.New("wrong password").Error(), http.StatusBadRequest)
		}
	}

}
