package repoi
import (
	"context"

	"github.com/jasurxaydarov/todo_app_kafka_server/models"
)
type UserRepoI interface{
	
	CreateUser(ctx context.Context, user models.User)error
	GetUSerById(ctx context.Context,id models.GetById)(*models.GetUserResp, error)

}