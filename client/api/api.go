package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/todo_app_kafka_client/api/handlers"
)


func Api(){

	h:=handlers.NewHandlers()
	engine:=gin.Default()

	api:=engine.Group("/api")


	api.POST("/create-user",h.CreateUser)
	api.POST("/get-user",h.GetUser)


	engine.Run()
}