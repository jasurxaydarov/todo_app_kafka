package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/todo_app_kafka_client/kafka"
	"github.com/jasurxaydarov/todo_app_kafka_client/models"
)

type handlers struct{

}

func NewHandlers()*handlers{

	return &handlers{}

}


func (h *handlers)CreateUser(ctx *gin.Context){

	var req models.User

	ctx.BindJSON(&req)

	err:=kafka.CreatUser(req)

	if err!=nil{
		log.Println(err)
		return
	}

	ctx.JSON(201,"created")
	
}

func (h *handlers)GetUser(ctx *gin.Context){

	var req models.GetById

	ctx.BindJSON(&req)

	err:=kafka.GetUserReq(req)

	if err!=nil{
		log.Println(err)
		return
	}

	resp:=kafka.GetUserResp("get-user-resp")
	ctx.JSON(200,resp)
	
}