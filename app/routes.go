package app

import "github.com/Ferinugroho/estore/app/controllers"

func (server *Server) InitializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
