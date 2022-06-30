package app

import (
	"memorable/internal/app/routers"

	"net/http"
	"time"
)

type App struct {
	server *http.Server
}

func InitApp() *App {
	return &App{}
}

func (m *App) InitHttpServer() {
	router := routers.InitRouters()
	m.server = &http.Server{
		Addr: ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func (m *App) Start() error {
	return m.server.ListenAndServe()
}