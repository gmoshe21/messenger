package models


type User struct {
	Uid 		string `json:"uid"`
	Name 		string `json:"name" validate:"required"`
	Mail 		string `json:"email" validate:"required"`
	Password 	string `json:"password" validate:"required"`
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

type LoginUserInput struct {
	Email    string `json:"email" bindinig:"required"`
	Password string `json:"password" binding:"required"`
}

type RegRec struct {
	Uid    string `json:"message" bindinig:"required"`
	Mes string `json:"status" binding:"required"`
}

type OTPInput struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type OTPRec struct {
	Base32 string `json:"base32"`
	OtpauthUrl string `json:"otpauth_url"`
}