package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func IsAlive(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Header)
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func Serve(port string) {
	router := mux.NewRouter()
	router.HandleFunc("/api/ping", IsAlive)
	addr := strings.Join([]string{"127.0.0.1", port}, ":")
	srv := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("start serving at %s \n", addr)
	log.Fatal(srv.ListenAndServe())
}
