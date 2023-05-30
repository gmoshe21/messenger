package control

import (
	"Messege/internal/models"
	"context"
)

type UseCase interface {
	CreateUser(ctx context.Context, params models.User) error
	FriendRequest(ctx context.Context, params models.FriendRequest) error
	CreateCommunication(ctx context.Context, params models.Communication) error
	CreateMessege(ctx context.Context, params models.Messege) error

	GetUsers(ctx context.Context) (result []byte, err error)
	GetFriends(ctx context.Context, uid string) (result []byte, err error)
	GetMesseges(ctx context.Context, author string, recipient string) (result []byte, err error)

}
