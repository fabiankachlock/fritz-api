package main

import (
	"fmt"
	"os"

	"github.com/fabiankachlock/fritz-api"
	"github.com/fabiankachlock/fritz-api/pkg/request"
	"github.com/fabiankachlock/fritz-api/pkg/response"
)

const (
	Username = "box"
	Password = "admin123"
)

func main() {
	client := fritz.NewClient("http://127.0.0.1:4000")
	err := client.Login(Username, Password)
	if err != nil {
		fmt.Printf("cant log in: %s\n", err)
		os.Exit(1)
	}
	session, err := client.GetSession()
	if err != nil {
		fmt.Printf("cant get session: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("session: %+v\n", session)

	body, err := client.GetData(request.NetworkUsageRequest)
	if err != nil {
		fmt.Printf("cant request data: %s\n", err)
		os.Exit(1)
	}

	// os.WriteFile("response.json", body, 0644)

	resp, err := response.UnmarshalCustomAs[response.NetCnt](body)
	if err != nil {
		fmt.Printf("cant unmarshal data: %s\n", err)
		os.Exit(1)
	}
	// marshalled, _ := json.Marshal(resp)
	// os.WriteFile("my.json", marshalled, 0644)

	// fmt.Println("connected devices:")
	// for _, d := range resp.Data.Devices {
	// 	if !d.OwnEntry {
	// 		fmt.Printf("Device: %s; connection type: %s; is self: %t\n", d.NameInfo.Name, d.ConnInfo.Kind, d.OwnClientDevice)
	// 	}
	// }

	err = client.Logout()
	if err != nil {
		fmt.Printf("cant logout: %s\n", err)
		os.Exit(1)
	}
}
