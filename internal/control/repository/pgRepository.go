package repository

import (
	"Messege/internal/control"
	"Messege/internal/models"
	"context"

	"github.com/jmoiron/sqlx"
	// "github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type controlRepo struct {
	db *sqlx.DB
}

func NewControlRepository(db *sqlx.DB) control.Repository {
	return &controlRepo{db: db}
}

func (c *controlRepo) CreateUser(ctx context.Context, params models.User) error {
	_, err := c.db.ExecContext(
		ctx,
		queryCreateUser,
		params.Uid,
		params.Name,
		params.Lastname,
		params.Number,
		params.Mail,
	)

	if err != nil {
		return errors.Wrap(err, "controlRepo.CreateUser.ExecContext()")
	}
	return nil
}

func (c *controlRepo) CreateFriendRequest(ctx context.Context, params models.FriendRequest) error {
	_, err := c.db.ExecContext(
		ctx,
		queryCreateFriendRequest,
		params.User1,
		params.User2,
	)

	if err != nil {
		return errors.Wrap(err, "controlRepo.CreateFriendRequest.ExecContext()")
	}

	return nil
}

func (c *controlRepo) CreateCommunication(ctx context.Context, user string, room string) error {
	_, err := c.db.ExecContext(
		ctx,
		queryCreateCommunication,
		user,
		room,
	)

	if err != nil {
		return errors.Wrap(err, "controlRepo.CreateCommunication.ExecContext()")
	}

	return nil
}

func (c *controlRepo) DeleteFriendRequest(ctx context.Context, params models.FriendRequest) error {
	_, err := c.db.ExecContext(
		ctx,
		queryDeleteFriendRequest,
		params.User1,
		params.User2,
	)

	if err != nil {
		return errors.Wrap(err, "controlRepo.DeleteFriendRequest.ExecContext()")
	}

	return nil
}

func (c *controlRepo) CreateMessege(ctx context.Context, params models.Messege) error {
	_, err := c.db.ExecContext(
		ctx,
		queryCreateMessege,
		params.Author,
		params.Recipient,
		params.Data,
		params.Time,
	)

	if err != nil {
		return errors.Wrap(err, "controlRepo.CreateMessege.ExecContext()")
	}

	return nil
}

func (c *controlRepo) GetUsers(ctx context.Context) (result []byte, err error) {
	err = c.db.GetContext(ctx, &result, queryGetUsers)

	if err != nil {
		return nil, errors.Wrap(err, "controlRepo.GetUsers.GetContext()")
	}

	return result, nil
}

func (c *controlRepo) GetFriends(ctx context.Context, uid string) (result []byte, err error) {
	err = c.db.GetContext(ctx, &result, queryGetFriends, uid)

	if err != nil {
		return nil, errors.Wrap(err, "controlRepo.GetFriends.GetContext()")
	}

	return result, nil
}

func (c *controlRepo) GetMesseges(ctx context.Context, author string, recipient string) (result []byte, err error) {
	err = c.db.GetContext(ctx, &result, queryGetMesseges, author, recipient)

	if err != nil {
		return nil, errors.Wrap(err, "controlRepo.GetMesseges.GetContext()")
	}

	return result, nil
}

func (c *controlRepo) GetFriendRequest(ctx context.Context, user string) (result []byte, err error) {
	err = c.db.GetContext(ctx, &result, queryGetFriendRequest, user)

	if err != nil {
		return nil, errors.Wrap(err, "controlRepo.GetFriendRequest.GetContext()")
	}

	return result, nil
}

func (c *controlRepo) GetKey(ctx context.Context, author string, recipient string) (result []byte, err error) {
	err = c.db.GetContext(ctx, &result, queryGetKey, author, recipient)

	if err != nil {
		return nil, errors.Wrap(err, "controlRepo.GetKey.GetContext()")
	}

	return result, nil
}

func (c *controlRepo) DeleteKey(ctx context.Context, author string, recipient string) error {
	_, err := c.db.ExecContext(ctx, queryDeleteKey, author, recipient)

	if err != nil {
		return errors.Wrap(err, "controlRepo.queryDeleteKey.ExecContext()")
	}

	return nil
}

