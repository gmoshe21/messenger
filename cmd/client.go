package main

import "fmt"

type Client struct {
	uid string
	friends []Friend
}

type Friend struct {
	uid string
	key string
}

type User struct {
	Uid 		string `json:"uid" validate:"required"`
	Name 		string `json:"name" validate:"required"`
	Lastname 	string `json:"lastname" validate:"required"`
	Number 		string `json:"number" validate:"required"`
	Mail 		string `json:"mail" validate:"required"`
}

func addUser(data User) {

}

func GetUsers() {
	//todo

	fmt.Println("1.Добавить 2.Назад")
}

func GetFriends() {
	

	fmt.Println("1.Написать 2.Назад")
}

func GetFriendRequest() {


	fmt.Println("1.Добавить 2.Назад")
}

func main() {
	var data User
	fmt.Println("Регистрация\n Имя:")
	fmt.Scan(&data.Name)
	fmt.Println("Фамилия:\n")
	fmt.Scan(&data.Name)
	fmt.Println("Номер телефона:\n")
	fmt.Scan(&data.Name)
	fmt.Println("Почта:\n")
	fmt.Scan(&data.Name)

	addUser(data)

	for {
		var command string
		fmt.Println("1.Пользователи 2.Сообщения 3.Запросы в друзья\n")
		fmt.Scan(&command)
		switch command {
			case "1":
				GetUsers()
				continue
			case "2":
				GetFriends()
				continue
			case "3":
				GetFriendRequest()
				continue
		}
	}
}
