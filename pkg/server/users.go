package server

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	"github.com/hoomanist/allegro-server/pkg/database"
)

func Hash(bv []byte) string {
	hasher := sha256.New()
	hasher.Write(bv)
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func (s *server) NewUser(w http.ResponseWriter, r *http.Request) {
	var m database.User
	json.NewDecoder(r.Body).Decode(&m)
	m.CreationDate = time.Now()
	m.Password = Hash([]byte(m.Password))
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
