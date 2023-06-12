package models


type User struct {
	Uid 		string `json:"uid" validate:"required"`
	Name 		string `json:"name" validate:"required"`
	Lastname 	string `json:"lastname" validate:"required"`
	Number 		string `json:"number" validate:"required"`
	Mail 		string `json:"mail" validate:"required"`
}

type FriendRequest struct{
	User1 		string `json:"user1" validate:"required"`
	User2 		string `json:"user2" validate:"required"`
}

type Communication struct {
	User1 		string `json:"user1" validate:"required"`
	User2 		string `json:"user2" validate:"required"`
}

type Messege struct {
	Author 		string `json:"author" validate:"required"`
    Recipient 	string `json:"recipient"`
    Data 		string `json:"data" validate:"required"`
    Time 		string `json:"time" validate:"required"`
}