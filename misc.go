package main

import (
  "net/http"
  "log"
  "encoding/json"
)

func (app *App) HandlePing() http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    response, _ := json.Marshal(map[string]interface{}{
      "message": "pong",
    })
    error_res, _ := json.Marshal(map[string]interface{}{
      "message": "db unavalible",
    })
    err := app.db.Ping()
    if  err == nil {
      w.Write([]byte(response))
      return
    }
    log.Fatal(err)
    w.Write([]byte(error_res))
  }
}
