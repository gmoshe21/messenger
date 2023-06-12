package usecase

import (
	"Messege/config"
	"Messege/internal/control"
	"Messege/internal/models"
	"context"
	"log"
	"sync"

	"github.com/google/uuid"
)

type controlUC struct {
	cfg         *config.Config
	mu          *sync.RWMutex
	controlRepo control.Repository
}

func NewControlUseCase(cfg *config.Config, controlRepo control.Repository) control.UseCase {
	return &controlUC{cfg: cfg, mu: &sync.RWMutex{}, controlRepo: controlRepo}
}

func (c *controlUC) CreateUser(ctx context.Context, params models.User) error {
	return c.controlRepo.CreateUser(ctx, params)
}

func (c *controlUC) FriendRequest(ctx context.Context, params models.FriendRequest) error {
	return c.controlRepo.CreateFriendRequest(ctx, params)
}

func (c *controlUC) CreateCommunication(ctx context.Context, params models.Communication) error {
	room := uuid.New().String()

	err := c.controlRepo.DeleteFriendRequest(ctx, models.FriendRequest(params))
	if err != nil {
		return err
	}

	err = c.controlRepo.CreateCommunication(ctx, params.User1, room)
	if err != nil {
		return err
	}

	return c.controlRepo.CreateCommunication(ctx, params.User2, room)
}

func (c *controlUC) CreateMessege(ctx context.Context, params models.Messege) error {
	return c.controlRepo.CreateMessege(ctx, params)
}

func (c *controlUC) GetUsers(ctx context.Context) (result []byte, err error) {
	return c.controlRepo.GetUsers(ctx)
}

func (c *controlUC) GetFriends(ctx context.Context, uid string) (result []byte, err error) {
	return c.controlRepo.GetFriends(ctx, uid)
}

func (c *controlUC) GetMesseges(ctx context.Context, author string, recipient string) (result []byte, err error) {
	return c.controlRepo.GetMesseges(ctx, author, recipient)
}

func (c *controlUC) GetFriendRequest(ctx context.Context, user string) (result []byte, err error) {
	return c.controlRepo.GetFriendRequest(ctx, user)
}

func (c *controlUC) GetKey(ctx context.Context, author string, recipient string) (result []byte, err error) {
	result, err = c.controlRepo.GetKey(ctx, author, recipient)
	if err != nil {
		return nil, err
	}
	log.Println(1)
	err = c.controlRepo.DeleteKey(ctx, author, recipient)
	if err != nil {
		return nil, err
	}

	return result, nil
}
