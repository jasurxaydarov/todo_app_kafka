package main

import (
	"fmt"

	"github.com/jasurxaydarov/todo_app_kafka_server/config"
	"github.com/jasurxaydarov/todo_app_kafka_server/kafka"
	"github.com/jasurxaydarov/todo_app_kafka_server/pkg/db"
)

func main(){

	cfg:=config.Load()

	conn,err:=db.ConnectDB(cfg.PgConfig)

	if err!=nil{
		fmt.Println(err)
		return
	}

	fmt.Println(conn)

	kafka_1.RunConsumers(conn)

	select{}
}