package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/todo_app_kafka_server/storage/postgres"
	"github.com/jasurxaydarov/todo_app_kafka_server/storage/repoi"
)

type storage struct{

	userRepo repoi.UserRepoI
}

type StorageI interface{

	GetUserRepo()repoi.UserRepoI

}


func NewStorage(con *pgx.Conn)StorageI{

	return &storage{
		userRepo: postgres.NewUserRepo(con),
	}
}


func (s *storage)GetUserRepo()repoi.UserRepoI{

	return s.userRepo

}





