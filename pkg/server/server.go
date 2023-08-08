package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/ini.v1"
)

type server struct {
	SqlCfg *ini.Section
	Key    string
}

func (s *server) IsAlive(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func Serve(cfg *ini.File) {
	router := mux.NewRouter()
	s := server{
		SqlCfg: cfg.Section("DB"),
		Key:    cfg.Section("").Key("encryption_key").String(),
	}
	router.HandleFunc("/api/ping", s.IsAlive)
	router.HandleFunc("/api/q/composers", s.ListComposers)
	router.HandleFunc("/api/new/composer", s.NewComposer).Methods("POST")
	router.HandleFunc("/upload", s.FileUpload).Methods("POST")
	router.HandleFunc("/api/new/user", s.NewUser).Methods("POST")
	router.HandleFunc("/api/q/users", s.GetUsers)
	router.HandleFunc("/login", s.Login).Methods("POST")
	addr := strings.Join([]string{"127.0.0.1", cfg.Section("").Key("port").String()}, ":")
	srv := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("start serving at %s \n", addr)
	log.Fatal(srv.ListenAndServe())
}
