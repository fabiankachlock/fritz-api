package main

import (
	"fmt"
	"os"

	"github.com/fabiankachlock/fritz-api/pkg/response"
)

const (
	Username = "box"
	Password = "admin123"
)

// func main() {
// 	client := api.NewClient("http://127.0.0.1:4000")
// 	err := client.Login(Username, Password)
// 	fmt.Println(err)
// 	session, _ := client.GetSession()
// 	fmt.Println(session)

// 	body, err := client.RequestData(request.HomeNetRequest)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	os.WriteFile("response.json", body, 0644)
// 	resp, err := response.UnmarshalAs[response.HomeNet](body)
// 	fmt.Println(resp, err)

// 	client.Logout()
// }

type A struct {
	B string `json:"b"`
}

type B struct {
	A
	Other string `json:"a"`
}

func main() {
	bytes, _ := os.ReadFile("response.json")
	resp, _ := response.UnmarshalAs[response.HomeNet](bytes)
	// fmt.Println(resp)
	for _, d := range resp.Data.Devices {
		fmt.Printf("Device: %s %s\n", d.Nameinfo.Name, d.Wlaninfo[0].Text)
	}
}
