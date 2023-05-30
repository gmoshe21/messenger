package control

import (
	"Messege/internal/models"
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context, params models.User) error
	CreateFriendRequest(ctx context.Context, params models.FriendRequest) error
	CreateCommunication(ctx context.Context, user string, room string) error
	CreateMessege(ctx context.Context, params models.Messege) error

	DeleteFriendRequest(ctx context.Context, params models.FriendRequest) error

	GetUsers(ctx context.Context) (result []byte, err error)
	GetFriends(ctx context.Context, uid string) (result []byte, err error)
	GetMesseges(ctx context.Context, author string, recipient string) (result []byte, err error)
}
