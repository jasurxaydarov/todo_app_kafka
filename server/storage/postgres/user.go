package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/todo_app_kafka_server/models"
	"github.com/jasurxaydarov/todo_app_kafka_server/storage/repoi"
)

type userRepo struct {
	db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) repoi.UserRepoI {

	return &userRepo{db: db}
}

func (u *userRepo) CreateUser(ctx context.Context, user models.User) error {

	query := `
	INSERT INTO 
		users (
			user_name, 
			password, 
			created_at
		) VALUES (
			 $1, $2, $3
		)`
	_, err := u.db.Exec(
		ctx, query,
		user.Username,
		user.Password,
		user.CreatedAt,
	)

	if err != nil {

		log.Println("error on CreateUser", err)
		return err
	}

	return nil
}
func (u *userRepo) GetUSerById(ctx context.Context, id models.GetById) (*models.GetUserResp, error) {
	var user models.GetUserResp


	
	query := `
		SELECT 
			user_name,
			password
		FROM
			users
		WHERE
			user_id=$1 
	`

	err := u.db.QueryRow(
		ctx, query,
		id.Id,
	).Scan(
		&user.Username,
		&user.Password,
	)

	if err != nil {

		log.Println("err on aaaaaaa GetUSerById", err)
		return nil, err

	}

	return &user, nil
}
