package main

import (
	"Messege/internal/models"
	"Messege/pkg/crypto"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"sort"

	"fmt"
	"os"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

var client *fasthttp.Client
var clientData Client
var Rsa crypto.RsaCrypto

type Client struct {
	Uid string `json:"uid" validate:"required"`
	Friends []Friend `json:"friends"`
	SecKey string `json:"SecKey" validate:"required"`
	PubKey string `json:"PubKey" validate:""`
}

type Friend struct {
	Uid string `json:"uid" validate:"required"`
	Key string `json:"key" validate:"required"`
}

func addUser(data models.User) {
	reqTimeout := time.Duration(100) * time.Millisecond

	reqUserBytes, _ := json.Marshal(data)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://localhost:8080/control/create_user")
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	req.SetBodyRaw(reqUserBytes)

	resp := fasthttp.AcquireResponse()
	err := client.DoTimeout(req, resp, reqTimeout)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if err != nil {
		errName, known := httpConnError(err)
		if known {
			fmt.Fprintf(os.Stderr, "WARN conn error: %v\n", errName)
		} else {
			fmt.Fprintf(os.Stderr, "ERR conn failure: %v %v\n", errName, err)
		}

		return
	}

	statusCode := resp.StatusCode()
	respBody := resp.Body()
	fmt.Printf("DEBUG Response: %s\n", respBody)

	if statusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "ERR invalid HTTP response code: %d\n", statusCode)

		return
	}

	// respEntity := &Entity{}
	// err = json.Unmarshal(respBody, respEntity)
	// if err == nil || errors.Is(err, io.EOF) {
	// 	fmt.Printf("DEBUG Parsed Response: %v\n", respEntity)
	// } else {
	// 	fmt.Fprintf(os.Stderr, "ERR failed to parse response: %v\n", err)
	// }
}

func httpConnError(err error) (string, bool) {
	var (
		errName string
		known   = true
	)

	switch {
	case errors.Is(err, fasthttp.ErrTimeout):
		errName = "timeout"
	case errors.Is(err, fasthttp.ErrNoFreeConns):
		errName = "conn_limit"
	case errors.Is(err, fasthttp.ErrConnectionClosed):
		errName = "conn_close"
	case reflect.TypeOf(err).String() == "*net.OpError":
		errName = "timeout"
	default:
		known = false
	}

	return errName, known
}

func GetUsers() {
	var users []models.User

	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://localhost:8080/control/get_users")
	req.Header.SetMethod(fasthttp.MethodGet)
	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	if err == nil {
		err = json.Unmarshal(resp.Body(), &users)
			if err == nil {
				ind := 1
				tmp := 0
				for i := 0; i < len(users); i++ {
					if users[i].Uid != clientData.Uid {
						fmt.Printf("%d. %s %s\n", ind, users[i].Name, users[i].Lastname)
						ind++
					} else {
						tmp = i
					}
				}
				users = append(users[:tmp], users[tmp+1:]...)
			} else {
				fmt.Println(err)
			}
	} else {
		fmt.Fprintf(os.Stderr, "ERR Connection error: %v\n", err)
	}
	fasthttp.ReleaseResponse(resp)

	for {
		var command int
		fmt.Print("\n*************************************************\n* 1.Добавить 2.Назад                            *\n*************************************************\n\nКоманда: ")
		fmt.Scan(&command)
		fmt.Println()

		if command == 2 {
			return
		}
		fmt.Print("\nВведите номер пользователя: ")
		fmt.Scan(&command)

		data := models.FriendRequest {
			User1: clientData.Uid,
			User2: users[command-1].Uid,
		}
		fmt.Println(users[command-1].Uid)
		friendRequest(data)
		fmt.Println("\nЗапрос в друзья отправлен")
	}
}

func friendRequest(data models.FriendRequest) {
	reqTimeout := time.Duration(100) * time.Millisecond

	reqUserBytes, _ := json.Marshal(data)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://localhost:8080/control/friend_request")
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	req.SetBodyRaw(reqUserBytes)

	resp := fasthttp.AcquireResponse()
	err := client.DoTimeout(req, resp, reqTimeout)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if err != nil {
		errName, known := httpConnError(err)
		if known {
			fmt.Fprintf(os.Stderr, "WARN conn error: %v\n", errName)
		} else {
			fmt.Fprintf(os.Stderr, "ERR conn failure: %v %v\n", errName, err)
		}

		return
	}

	statusCode := resp.StatusCode()
	respBody := resp.Body()
	fmt.Printf("DEBUG Response: %s\n", respBody)

	if statusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "ERR invalid HTTP response code: %d\n", statusCode)

		return
	}
}

func getKey(author string, recipient string) {
	var messeges []models.Messege
fmt.Println("getkey")

	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://localhost:8080/control/get_key")
	req.Header.SetMethod(fasthttp.MethodGet)
	req.URI().QueryArgs().Add("author", author)
	req.URI().QueryArgs().Add("recipient", recipient)
	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	if err == nil {
		fmt.Println("unmar")
		err = json.Unmarshal(resp.Body(), &messeges)
		if err == nil {
			fmt.Println(clientData.Uid, clientData.Friends, author)
				for i := 0; i < len(clientData.Friends); i++ {
					if clientData.Friends[i].Uid == author {
						clientData.Friends[i].Key = messeges[0].Data
						break
					}
				}
				fmt.Println("writetofile")
				writeToFile()
			} else {
				fmt.Println(err)
			}
	} else {
		fmt.Fprintf(os.Stderr, "ERR Connection error: %v\n", err)
	}
	fasthttp.ReleaseResponse(resp)
}

func getFriends() {
	var users []models.User

	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://localhost:8080/control/get_friends")
	req.Header.SetMethod(fasthttp.MethodGet)
	req.URI().QueryArgs().Add("uid", clientData.Uid)
	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	if err == nil {
		err = json.Unmarshal(resp.Body(), &users)
			if err == nil {
				for i := 0; i < len(users); i++ {
					if users[i].Uid != clientData.Uid {
						fmt.Printf("%d. %s %s\n", i+1, users[i].Name, users[i].Lastname)
					}
				}
			} else {
				if err.Error() == "unexpected end of JSON input" {
					fmt.Println("\n---Список пуст---\n")
					return
				}
				fmt.Println(err)
			}
	} else {
		fmt.Fprintf(os.Stderr, "ERR Connection error: %v\n", err)
	}
	fasthttp.ReleaseResponse(resp)

	for {
		var command int
		fmt.Println("\n*************************************************\n* 1.Написать 2.Назад                            *\n*************************************************\n")
		fmt.Print("Команда: ")
		fmt.Scan(&command)
		fmt.Println()

		if command == 2 {
			return
		}
		fmt.Print("Введите номер пользователя: ")
		fmt.Scan(&command)

		data := models.Messege {
			Author: clientData.Uid,
			Recipient : users[command-1].Uid,
			Data : "",
			Time: "",

		}

		if len(clientData.Friends) == 0 {
			data.Time = time.Now().Format(time.RFC3339)
			data.Data = clientData.SecKey
			addMessege(data)
			clientData.Friends = append(clientData.Friends, Friend{Uid: users[command-1].Uid})
		}

		for i := 0; i < len(clientData.Friends); i++ {
			fmt.Println(2)
			if clientData.Friends[i].Uid == users[command-1].Uid && clientData.Friends[i].Key == "" {
				getKey(clientData.Friends[i].Uid, clientData.Uid)
			} else if clientData.Friends[i].Uid == users[command-1].Uid && len(clientData.Friends[i].Key) != 0 {
				communication(data, clientData.Friends[i].Key)
				break
			} else {
				data.Time = time.Now().Format(time.RFC3339)
				data.Data = clientData.SecKey
				addMessege(data)
				clientData.Friends = append(clientData.Friends, Friend{Uid: users[command-1].Uid})
			}
		}

		// fmt.Println("\nЗапрос в друзья отправлен")
	}

}

func communication(data models.Messege, key string) {
	var err error
	fmt.Println("1. - Вернуться")
	getMessege(data.Author, data.Recipient, key)
	for {
		tmp := ""
		fmt.Scan(&tmp)
		if tmp == "1." {
			return
		}
		data.Data, err = Rsa.Encrypt(tmp, clientData.PubKey)
		if err != nil {
			fmt.Println(err)
			break
		}

		data.Time = time.Now().Format(time.RFC3339)
		addMessege(data)
	}

}

func getMessege(author string, recipient string, key string) {
	var messeges []models.Messege


	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://localhost:8080/control/get_messeges")
	req.Header.SetMethod(fasthttp.MethodGet)
	req.URI().QueryArgs().Add("author", author)
	req.URI().QueryArgs().Add("recipient", recipient)
	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	if err == nil {
		err = json.Unmarshal(resp.Body(), &messeges)
		if err == nil {
				sort.SliceStable(messeges, func(i, j int) bool {
					return messeges[i].Time < messeges[j].Time
				})
				for i := 0; i < len(messeges); i++ {
					fmt.Println(key)
					str, err := Rsa.Decrypt(messeges[i].Data, key, "go")
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println("end")
					fmt.Println("(", messeges[i].Time, ")", " ", messeges[i].Author, ": ", str)
				}
			} else {
				fmt.Println(err)
			}
	} else {
		fmt.Fprintf(os.Stderr, "ERR Connection error: %v\n", err)
	}
	fasthttp.ReleaseResponse(resp)
}

func addMessege(data models.Messege) {
	reqTimeout := time.Duration(100) * time.Millisecond

	reqUserBytes, _ := json.Marshal(data)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://localhost:8080/control/create_messege")
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	req.SetBodyRaw(reqUserBytes)

	resp := fasthttp.AcquireResponse()
	err := client.DoTimeout(req, resp, reqTimeout)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if err != nil {
		errName, known := httpConnError(err)
		if known {
			fmt.Fprintf(os.Stderr, "WARN conn error: %v\n", errName)
		} else {
			fmt.Fprintf(os.Stderr, "ERR conn failure: %v %v\n", errName, err)
		}

		return
	}

	statusCode := resp.StatusCode()
	respBody := resp.Body()
	fmt.Printf("DEBUG Response: %s\n", respBody)

	if statusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "ERR invalid HTTP response code: %d\n", statusCode)

		return
	}
}

func getFriendRequest() {
	var users []models.User

	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://localhost:8080/control/get_friend_request")
	req.Header.SetMethod(fasthttp.MethodGet)
	req.URI().QueryArgs().Add("user", clientData.Uid)
	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	if err == nil {
		err = json.Unmarshal(resp.Body(), &users)
			if err == nil {
				for i := 0; i < len(users); i++ {
					if users[i].Uid != clientData.Uid {
						fmt.Printf("%d. %s %s\n", i+1, users[i].Name, users[i].Lastname)
					}
				}
			} else {
				if err.Error() == "unexpected end of JSON input" {
					fmt.Println("\n---Список пуст---\n")
					return
				}
				fmt.Println(err)
			}
	} else {
		fmt.Fprintf(os.Stderr, "ERR Connection error: %v\n", err)
	}
	fasthttp.ReleaseResponse(resp)

	for {
		var command int
		fmt.Println("\n*************************************************\n* 1.Добавить 2.Назад                            *\n*************************************************\n")
		fmt.Print("Команда: ")
		fmt.Scan(&command)
		fmt.Println()

		if command == 2 {
			return
		}
		fmt.Print("Введите номер пользователя: ")
		fmt.Scan(&command)

		data := models.FriendRequest {
			User2: users[command-1].Uid,
			User1: clientData.Uid,
		}

		// fmt.Println(clientData.Keys.PublicKey.N.String())
		createCommunication(data)

		clientData.Friends = append(clientData.Friends, Friend {
			Uid: users[command-1].Uid,
		})
		fmt.Println(clientData)
		addMessege(models.Messege{
			Author: clientData.Uid,
			Recipient: users[command-1].Uid,
			Data: clientData.SecKey,
			Time:  time.Now().Format(time.RFC3339),
		})

		fmt.Println("\nЗапрос в друзья отправлен")
	}
}

func createCommunication(data models.FriendRequest) {
	reqTimeout := time.Duration(100) * time.Millisecond

	reqUserBytes, _ := json.Marshal(data)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://localhost:8080/control/create_communication")
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	req.SetBodyRaw(reqUserBytes)

	resp := fasthttp.AcquireResponse()
	err := client.DoTimeout(req, resp, reqTimeout)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if err != nil {
		errName, known := httpConnError(err)
		if known {
			fmt.Fprintf(os.Stderr, "WARN conn error: %v\n", errName)
		} else {
			fmt.Fprintf(os.Stderr, "ERR conn failure: %v %v\n", errName, err)
		}

		return
	}

	statusCode := resp.StatusCode()
	respBody := resp.Body()
	fmt.Printf("DEBUG Response: %s\n", respBody)

	if statusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "ERR invalid HTTP response code: %d\n", statusCode)

		return
	}
}

func commands() {
	for {
		var command string
		fmt.Println("\n*************************************************\n* 1.Пользователи 2.Сообщения 3.Запросы в друзья *\n*************************************************")
		fmt.Print("\nКоманда: ")
		fmt.Scan(&command)
		switch command {
			case "1":
				GetUsers()
				continue
			case "2":
				getFriends()
				continue
			case "3":
				getFriendRequest()
				continue
		}
	}
}

func registers() {
	var data models.User
	var err error

	fmt.Println("Регистрация\n Имя:")
	fmt.Scan(&data.Name)
	fmt.Println("Фамилия:\n")
	fmt.Scan(&data.Lastname)
	fmt.Println("Номер телефона:\n")
	fmt.Scan(&data.Number)
	fmt.Println("Почта:\n")
	fmt.Scan(&data.Mail)

	data.Uid = uuid.New().String()
	addUser(data)
	clientData.Uid = data.Uid


	// Keys, _ := rsa.GenerateKey(rand.Reader, 256)
	// Keys.B
	clientData.SecKey, clientData.PubKey, err = Rsa.GenerateKeyPair(2048)
	if err != nil {
		fmt.Println(err)
	}
	if "" == clientData.SecKey || "" == clientData.PubKey {
		fmt.Println("Error generating key pair")
	}

	writeToFile()
}

func writeToFile() {
	str, err := json.Marshal(clientData)
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.Create("Client.json")
    if err != nil{
        fmt.Println("Unable to create file:", err) 
        os.Exit(1) 
    }
    defer file.Close() 
    file.Write(str)
}

func checkRegist() bool {
	// file, err := os.Open("Client.json")
    // if err != nil{
    //     fmt.Println(err) 
    //     os.Exit(1) 
    // }
    // defer file.Close() 
     
    // data := make([]byte, 64)
     
    // for{
    //     _, err := file.Read(data)
    //     if err == io.EOF{   // если конец файла
    //         break           // выходим из цикла
    //     }
	// 	data = data[:len(data)-1]
    //     fmt.Print(string(data))
	// 	// fmt.Println()
    // }
	// err = json.Unmarshal(data, &clientData)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // fmt.Println(clientData)
	cfdFile, err := LoadConfig()
	if err != nil {
		fmt.Println("LoadConfig: %v", err)
	}
	ParseConfig(cfdFile)

	if err != nil {
		fmt.Println(err)
	}
	if len(clientData.Uid) == 0 {
		return false
	}

	return true
}

func main() {
	readTimeout, _ := time.ParseDuration("500ms")
	writeTimeout, _ := time.ParseDuration("500ms")
	maxIdleConnDuration, _ := time.ParseDuration("1h")
	client = &fasthttp.Client{
		ReadTimeout:                   readTimeout,
		WriteTimeout:                  writeTimeout,
		MaxIdleConnDuration:           maxIdleConnDuration,
		NoDefaultUserAgentHeader:      true, // Don't send: User-Agent: fasthttp
		DisableHeaderNamesNormalizing: true, // If you set the case on your headers correctly you can enable this
		DisablePathNormalizing:        true,
		// increase DNS cache time to an hour instead of default minute
		Dial: (&fasthttp.TCPDialer{
			Concurrency:      4096,
			DNSCacheDuration: time.Hour,
		}).Dial,
	}

	if checkRegist() == false {
		registers()
	}
	commands()
}

func sendGetRequest() {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://localhost:8080/control/get_users")
	req.Header.SetMethod(fasthttp.MethodGet)
	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	if err == nil {
		fmt.Printf("DEBUG Response: %s\n", resp.Body())
		} else {
		fmt.Fprintf(os.Stderr, "ERR Connection error: %v\n", err)
	}
	fasthttp.ReleaseResponse(resp)
}

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath(fmt.Sprintf("./"))
	v.SetConfigName("Client")
	v.SetConfigType("json")
	if err := v.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	return v, nil
}

func ParseConfig(v *viper.Viper) (){
	err := v.Unmarshal(&clientData)
	fmt.Println(clientData)
	if err != nil {
		fmt.Println("unable to decode into struct, %v", err)
	}
	err = validator.New().Struct(clientData)
	if err != nil {
		fmt.Println(err)
	}
}