package main

import (
	"fmt"

	"github.com/fabiankachlock/fritz-api/pkg/api"
)

const (
	DataUrl = "http://127.0.0.1:4000/data.lua"

	Username = "box"
	Password = "admin123"
)

func main() {
	client := api.NewClient("http://127.0.0.1:4000")
	err := client.Login(Username, Password)
	fmt.Println(err)
	session, _ := client.GetSession()
	fmt.Println(session)
	client.Logout()
}
