package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hoomanist/allegro-server/pkg/database"
)

func (s *server) NewComposer(w http.ResponseWriter, r *http.Request) {
	var m database.Composer
	fmt.Println(r.Body)
	json.NewDecoder(r.Body).Decode(&m)
	err := database.NewComposer(s.SqlCfg, &m)
	log.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "wasn't able to insert a new entry",
		})
	}

	json.NewEncoder(w).Encode(map[string]string{
		"status": "done",
	})
}

func (s *server) ListComposers(w http.ResponseWriter, r *http.Request) {
	composers, err := database.ListComposers(s.SqlCfg)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
	}
	j, _ := json.Marshal(composers)
	w.Write(j)
}
