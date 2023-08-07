package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/hoomanist/allegro-server/pkg/database"
)

func (s *server) NewUser(w http.ResponseWriter, r *http.Request) {
	var m database.User
	json.NewDecoder(r.Body).Decode(&m)
	m.CreationDate = time.Now()
	err := database.NewUser(s.SqlCfg, m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"status": "done",
	})
}

func (s *server) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := database.GetUsers(s.SqlCfg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	j, _ := json.Marshal(users)
	w.Write(j)
}
