package response

import "encoding/xml"

type SessionInfo struct {
	XMLName     xml.Name `xml:"SessionInfo"`
	SID         string   `xml:"SID"`           // SID is the current session id
	Challenge   string   `xml:"Challenge"`     // Challenge is the challenge string for login
	BlockTime   int      `xml:"BlockTime"`     // BlockTime is the time in seconds until the next login attempt is allowed
	Users       []User   `xml:"Users>User"`    // Users is a list of users
	RightName   []string `xml:"Rights>Name"`   // RightName is a list of rights
	RightAccess []int    `xml:"Rights>Access"` // RightAccess is a list of access for those rights
}

type User struct {
	Last int    `xml:"last,attr"`
	Name string `xml:",chardata"`
}

type Right struct {
	Name   string `xml:"Name"`
	Access int    `xml:"Access"`
}

func (s SessionInfo) GetRights() []Right {
	var rights []Right
	for i, name := range s.RightName {
		rights = append(rights, Right{Name: name, Access: s.RightAccess[i]})
	}
	return rights
}
