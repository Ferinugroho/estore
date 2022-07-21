package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

func (server *Server) Initialize(appCofig AppConfig) {
	fmt.Println("Welcome to " + appCofig.AppName)

	server.Router = mux.NewRouter()
	server.InitializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port $s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}

	appConfig.AppName = "eStore"
	appConfig.AppEnv = "Development"
	appConfig.AppPort = "9999"

	server.Initialize(appConfig)
	server.Run(":" + appConfig.AppPort)
}
