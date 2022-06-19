package main

import (
        "database/sql"
        "time"
        "embed"
        "net/http"
        "os"
        "log"
        "fmt"

        "github.com/pressly/goose/v3"
        "github.com/gorilla/mux"
        "github.com/gorilla/handlers"
        _ "github.com/go-sql-driver/mysql"
)

type App struct {
  router *mux.Router
  db *sql.DB
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

func DatabaseInit() (*sql.DB, error) {
  db, err := sql.Open("mysql", "hooman:hooman86@tcp(localhost:3306)/Allegro")
  if err != nil {
    panic(err)
  }
  db.SetConnMaxLifetime(time.Minute * 3)
  db.SetMaxOpenConns(10)
  db.SetMaxIdleConns(10)
  err = db.Ping()
  if err != nil {
    return nil, err
  }
  goose.SetBaseFS(embedMigrations)
  if err := goose.SetDialect("mysql"); err != nil {
      return nil, err
  }

  if err := goose.Up(db, "migrations"); err != nil {
      return nil, err
  }
  return db, nil

}

func NewApp() App {
  var err error
  app := App{}
  app.router = mux.NewRouter()
  app.db, err = DatabaseInit()
  if err == nil {
    return app
  }
  panic(err)
}

func main() {
  app := NewApp()
  app.SetupRoutes()
  router := handlers.LoggingHandler(os.Stdout, app.router)
  http.Handle("/", handlers.CompressHandler(router))
  fmt.Println("starting the server...")
  log.Fatal(http.ListenAndServe(":5000", nil))
}
