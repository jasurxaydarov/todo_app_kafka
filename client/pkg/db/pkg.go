package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/todo_app_kafka_client/config"
)

func ConnectDB(pgCfg config.PgConfig) (*pgx.Conn, error) {


	dbUrl:=fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		pgCfg.Username,
		pgCfg.Password,
		pgCfg.Host,
		pgCfg.Port,
		pgCfg.DatabaseName,
	)

	conn,err:=pgx.Connect(context.Background(),dbUrl)

	if err!= nil{

		log.Println("unable to connect with db ",err)

		return nil,err
	}


	return conn, nil
}