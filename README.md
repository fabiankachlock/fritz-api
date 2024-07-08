# fritz-api

An unofficial AVM FRITZ!Box WebUI scraping client written in Go.

![GitHub Release](https://img.shields.io/github/v/release/fabiankachlock/fritz-api?style=for-the-badge)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/fabiankachlock/fritz-api?style=for-the-badge)
![GitHub License](https://img.shields.io/github/license/fabiankachlock/fritz-api?style=for-the-badge)

This is a tool to scrape information from the FRITZ!Box web UI, like connection status or connected devices. 

It will only execute read only operations and is not able to change configuration or settings. For changing the configuration or missing critical information it's recommended to use a fully supported and officially documented protocol like `TR-064` or `IGD 2.0`. [See more](https://avm.de/service/schnittstellen/)

Since this client uses the internal `/data.lua` endpoint without any official documentation, it is not possible to provide full support or any guarantee that the data is correct. 

Due to responses having a polymorphic schema based on factors such as the type of device or its connection, the models returned by the client currently represent the union of all fields discovered during reverse engineering. This might change in the future. Consequently, not all fields will always be populated, and additional fields might exist but aren't picked up by the client. To circumvent such issues, the client provides methods that return the raw JSON mapping of the results, containing the full amount of information.


Docs: https://pkg.go.dev/github.com/fabiankachlock/fritz-api

## Login

The login works via a designed endpoint for third party applications at `/login_sid.lua` which is documented at: [https://avm.de/fileadmin/user_upload/Global/Service/Schnittstellen/AVM_Technical_Note_-_Session_ID_english_2021-05-03.pdf](https://avm.de/fileadmin/user_upload/Global/Service/Schnittstellen/AVM_Technical_Note_-_Session_ID_english_2021-05-03.pdf)

> [!CAUTION]
> The client currently support only solving `pbkdf2` based challenge sent by routers using Fritz!OS 7.25 or later

> [!CAUTION]
> The `/login_sid.lua` endpoint does not seem to work with the default user. 

In order to perform a login an username & password is needed. It's recommended to create a designed user with limited access rights to be used by this client at [http://fritz.box/#user](http://fritz.box/#user).

## Usage

Install the package:

`go get -u github.com/fabiankachlock/fritz-api`

See the example:

```go
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

	resp, err := response.UnmarshalCustomAs[response.NetCnt](body)
	if err != nil {
		fmt.Printf("cant unmarshal data: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("connected devices:")
	for _, d := range resp.Data.Devices {
        if !d.OwnEntry {
            fmt.Printf("Device: %s; connection type: %s; is self: %t\n", d.NameInfo.Name, d.ConnInfo.Kind, d.OwnClientDevice)
		}
	}

	err = client.Logout()
    if err != nil {
        fmt.Printf("cant unmarshal data: %s\n", err)
        os.Exit(1)
    }
}

```