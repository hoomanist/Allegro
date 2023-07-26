package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Header)
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func Serve(port string) {
	router := mux.NewRouter()
	router.HandleFunc("/", IndexPage)
	addr := strings.Join([]string{"127.0.0.1", port}, ":")
	srv := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
