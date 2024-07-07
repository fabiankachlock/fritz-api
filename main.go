package main

import (
	"fmt"
	"os"

	"github.com/fabiankachlock/fritz-api/pkg/api"
	"github.com/fabiankachlock/fritz-api/pkg/request"
	"github.com/fabiankachlock/fritz-api/pkg/response"
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

	body, err := client.RequestData(request.HomeNetRequest)
	if err != nil {
		fmt.Println(err)
	}
	os.WriteFile("response.json", body, 0644)
	resp, err := response.UnmarshalAs[response.HomeNet](body)
	fmt.Println(resp, err)

	client.Logout()
}
