package kafka_1

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/todo_app_kafka_server/models"
	"github.com/jasurxaydarov/todo_app_kafka_server/storage"
	"github.com/segmentio/kafka-go"
)


func CreateUser(conn *pgx.Conn,topic string) {

	storage:=storage.NewStorage(conn)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"0.0.0.0:9092"},
		GroupID: "github.com/jasurxaydarov/todo_app_kafka",
		Topic:   topic,
	})

	fmt.Println(fmt.Sprintln(topic," consumer is listenning"))
	for {

		var user models.User
		Kfmsg, err := reader.ReadMessage(context.Background())

		if err != nil {
			log.Println("err on ReadMessage", err)
			continue
		}

		err = json.Unmarshal(Kfmsg.Value, &user)

		if err != nil {
			log.Println("err on Unmarshal", err)
			continue
		}

		switch topic {
		case "create-user":

			storage.GetUserRepo().CreateUser(context.Background(),user)

			fmt.Println("new message:",user)


		}
	}

}

func GetUser(conn *pgx.Conn,topic string)models.GetUserResp {

	storage:=storage.NewStorage(conn)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"0.0.0.0:9092"},
		GroupID: "github.com/jasurxaydarov/todo_app_kafka",
		Topic:   topic,
	})

	fmt.Println(fmt.Sprintln(topic," consumer is listenning"))
	for {

		var req models.GetById
		Kfmsg, err := reader.ReadMessage(context.Background())

		if err != nil {
			log.Println("err on ReadMessage", err)
			continue
		}

		err = json.Unmarshal(Kfmsg.Value, &req.Id)

		if err != nil {
			log.Println("err on Unmarshal", err)
			continue
		}

		switch topic {
		case "get-user":

			resp,err:=storage.GetUserRepo().GetUSerById(context.Background(),req)

			if err!= nil{
				log.Println("err on GetUSerById ",err)
				continue
			}
			fmt.Println("new message:",req)


			GetUserResp(*resp)
			return *resp
		}
	}

}

func RunConsumers(conn *pgx.Conn) {

	go GetUser(conn,"get-user")

	go CreateUser(conn,"create-user")

}
