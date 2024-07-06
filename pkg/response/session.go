package response

import "encoding/xml"

type SessionInfo struct {
	XMLName   xml.Name `xml:"SessionInfo"`
	SID       string   `xml:"SID"`
	Challenge string   `xml:"Challenge"`
	BlockTime int      `xml:"BlockTime"`
	Users     []User   `xml:"Users>User"`
	Rights    []Right  `xml:"Rights>Name"`
}

type User struct {
	Last int    `xml:"last,attr"`
	Name string `xml:",chardata"`
}

type Right struct {
	Name   string `xml:"Name"`
	Access int    `xml:"Access"`
}
