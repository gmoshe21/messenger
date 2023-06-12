package http

import (
	"Messege/config"
	"Messege/internal/control"
	"Messege/internal/models"
	"Messege/pkg/reqvalidator"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
)

type controlHandlers struct { //Messege
	cfg       *config.Config
	controlUC control.UseCase
}

func NewControlHandlers(cfg *config.Config, controlUC control.UseCase) control.Handlers {
	return &controlHandlers{cfg: cfg, controlUC: controlUC}
}

func (ctrl *controlHandlers) CreateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {	
		var params models.User
		if err := reqvalidator.ReadRequest(c, &params); err != nil {
			log.Println("controlHandlers.CreateUser.ReadRequest", err)
			return err
		}

		err := ctrl.controlUC.CreateUser(context.Background(), params)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func (ctrl *controlHandlers) FriendRequest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params models.FriendRequest
		if err := reqvalidator.ReadRequest(c, &params); err != nil {
			log.Println("controlHandlers.FriendRequest.ReadRequest", err)
			return err
		}

		err := ctrl.controlUC.FriendRequest(context.Background(),  params)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func (ctrl *controlHandlers) CreateCommunication() fiber.Handler {
	return func(c *fiber.Ctx) error {	
		var params models.Communication
		if err := reqvalidator.ReadRequest(c, &params); err != nil {
			log.Println("controlHandlers.CreateCommunication.ReadRequest", err)
			return err
		}

		err := ctrl.controlUC.CreateCommunication(context.Background(), params)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func (ctrl *controlHandlers) CreateMessege() fiber.Handler {
	return func(c *fiber.Ctx) error {	
		var params models.Messege
		if err := reqvalidator.ReadRequest(c, &params); err != nil {
			log.Println("controlHandlers.Messege.ReadRequest", err)
			return err
		}

		err := ctrl.controlUC.CreateMessege(context.Background(), params)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func (ctrl *controlHandlers) GetUsers() fiber.Handler {
	return func(c *fiber.Ctx) error {	
		result, err := ctrl.controlUC.GetUsers(context.Background())
		if err != nil {
			return err
		}

		return c.Send(result)
	}
}

func (ctrl *controlHandlers) GetFriends() fiber.Handler {
	return func(c *fiber.Ctx) error {
		uid := c.Query("uid")
		result, err := ctrl.controlUC.GetFriends(context.Background(), uid)
		if err != nil { 
			return err
		}

		return c.Send(result)
	}
}

func (ctrl *controlHandlers) GetMesseges() fiber.Handler {
	return func(c *fiber.Ctx) error {
		author := c.Query("author")
		recipient := c.Query("recipient")
		result, err := ctrl.controlUC.GetMesseges(context.Background(), author, recipient)
		if err != nil {
			return err
		}

		return c.Send(result)
	}
}

func (ctrl *controlHandlers) GetFriendRequest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Query("user")
		result, err := ctrl.controlUC.GetFriendRequest(context.Background(), user)
		if err != nil {
			return err
		}

		return c.Send(result)
	}
}

func (ctrl *controlHandlers) GetKey() fiber.Handler {
	return func(c *fiber.Ctx) error {
		author := c.Query("author")
		recipient := c.Query("recipient")
		result, err := ctrl.controlUC.GetKey(context.Background(), author, recipient)
		if err != nil {
			return err
		}

		return c.Send(result)
	}
}
                