package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jasurxaydarov/todo_app_kafka_client/models"

	"github.com/segmentio/kafka-go"
)

func GetUserResp(topic string) models.GetUserResp {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"0.0.0.0:9092"},
		GroupID: "github.com/jasurxaydarov/todo_app_kafka",
		Topic:   topic,
	})

	fmt.Println(fmt.Sprintln(topic, " consumer is listenning"))
	for {

		var user models.GetUserResp
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
		case "get-user-resp":

			fmt.Println("new message:", user)

			return user

		}
	}

}
