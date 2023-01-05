package struct_demo

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-redis/redis/v8"
)

type Server struct {
	ServerName string `json:"name"`
	ServerHost string `json:"host"`
	ServerPort int `json:"port,8000"`
}

func TestStruct()  {
	//server := new(Server)
	server := Server{}
	server.ServerName = "localhost"
	server.ServerHost = "127.0.0.1"
	server.ServerPort = 8081

	obj, err := json.Marshal(server)
	if err != nil {

		fmt.Println("err: ", err.Error())
	}

	fmt.Println(string(obj))

	client := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})
	fmt.Println(client.PoolStats())
	fmt.Println(client)
}