package app

import (
	"fmt"
	"net/http"
	"time"
)

//Config hold the configuration for the server
type Config struct {
	Addr string  // PORT
}

//Application hold the Server details
type Application struct {
	Config Config
}

//Constructor for Config
func NewConfig(addr string) Config {
	return  Config{
		Addr: addr,
	}
} 

//Constructor for Application
func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
	}
}



func (app * Application) Run() error {
	server := &http.Server{
		Addr: app.Config.Addr,
		Handler: nil,  //Todo: setup chi router here
		ReadTimeout: 10 * time.Second, 
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 10 * time.Second,
	}

	fmt.Println("Starting server on port", app.Config.Addr)

	return  server.ListenAndServe()
}