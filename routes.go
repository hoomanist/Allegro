package main

func (app *App) SetupRoutes() {
  app.router.HandleFunc("/ping", app.HandlePing()).Methods("GET")
}
