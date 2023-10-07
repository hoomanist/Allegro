package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/hoomanist/allegro-server/pkg/auth"
	"gopkg.in/ini.v1"
)

type server struct {
	SqlCfg *ini.Section
	Key    string
}

func (s *server) IsAlive(w http.ResponseWriter, r *http.Request) {
    log.Println("Hallo from client")
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func Serve(cfg *ini.File) {
	router := mux.NewRouter()
    router.Use(mux.CORSMethodMiddleware(router))
	s := server{
		SqlCfg: cfg.Section("DB"),
		Key:    cfg.Section("").Key("encryption_key").String(),
	}
	router.HandleFunc("/api/ping", s.IsAlive)
	router.HandleFunc("/api/q/composers", s.ListComposers)
	router.HandleFunc("/api/new/composer", s.NewComposer).Methods("POST")
	router.HandleFunc("/upload", auth.Authorize(s.FileUpload, s.Key)).Methods("POST")
	router.HandleFunc("/api/new/user", s.NewUser).Methods("POST")
	router.HandleFunc("/api/q/users", s.GetUsers)
	router.HandleFunc("/login", s.Login).Methods("POST")
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
                                    http.FileServer(http.Dir("./uploads"))))


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
