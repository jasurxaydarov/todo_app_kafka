package kafka_1

import (
	"context"
	"encoding/json"
	"log"

	"github.com/jasurxaydarov/todo_app_kafka_server/models"
	"github.com/segmentio/kafka-go"
)

func SendMsg(topic string, msg []byte) error {

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"0.0.0.0:9092"},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: msg,
		},
	)

	if err != nil {
		log.Println("err on WriteMessages", err)
		return err
	}

	err = writer.Close()

	if err != nil {
		log.Println("err on Close", err)
		return err
	}

	return nil
}

func CreatUser(msg models.User) error {

	topic := "create-user"

	bytData, err := json.Marshal(msg)

	if err != nil {
		log.Println("err on Marshal", err)
		return err
	}

	SendMsg(topic, bytData)

	return nil
}

func GetUserResp(msg models.GetUserResp) error {

	topic := "get-user-resp"

	bytData, err := json.Marshal(msg)

	if err != nil {
		log.Println("err on Marshal", err)
		return err
	}

	SendMsg(topic, bytData)

	return nil
}
