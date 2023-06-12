package http

import (
	"Messege/internal/control"

	"github.com/gofiber/fiber/v2"
)

func MapAPIRoutes(group fiber.Router, h control.Handlers) {
	group.Get("/get_users", h.GetUsers())
	group.Get("/get_friends", h.GetFriends())
	group.Get("/get_messeges", h.GetMesseges())
	group.Get("/get_friend_request", h.GetFriendRequest())
	group.Get("/get_key", h.GetKey())

	group.Post("/create_user", h.CreateUser())
	group.Post("/friend_request", h.FriendRequest())
	group.Post("/create_communication", h.CreateCommunication())
	group.Post("/create_messege", h.CreateMessege())
}
