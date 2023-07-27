package server

import (
	"encoding/json"
	"net/http"

	"github.com/hoomanist/allegro-server/pkg/database"
)

func NewComposer(w http.ResponseWriter, r *http.Request) {

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
