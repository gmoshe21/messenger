package control

import (
	"github.com/gofiber/fiber/v2"
)

type Handlers interface {
	CreateUser() fiber.Handler
	FriendRequest() fiber.Handler
	CreateCommunication() fiber.Handler
	CreateMessege() fiber.Handler

	GetUsers() fiber.Handler
	GetFriends() fiber.Handler
	GetMesseges() fiber.Handler
}