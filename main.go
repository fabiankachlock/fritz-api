package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

const (
	LoginUrl = "http://127.0.0.1:4000//login_sid.lua?version=2"
	DataUrl  = "http://127.0.0.1:4000/data.lua"

	Username = "box"
	Password = "admin123"
)

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

func challenge(challenge string, pw string) string {
	parts := strings.Split(challenge, "$")
	iter1, _ := strconv.Atoi(parts[1])
	iter2, _ := strconv.Atoi(parts[3])
	salt1, _ := hex.DecodeString(parts[2])
	salt2, _ := hex.DecodeString(parts[4])
	hash1 := pbkdf2.Key([]byte(pw), salt1, iter1, 32, sha256.New)
	hash2 := pbkdf2.Key(hash1, salt2, iter2, 32, sha256.New)
	return parts[4] + "$" + hex.EncodeToString(hash2)
}

func sendLoginRequest(method string, body io.Reader) (SessionInfo, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, LoginUrl, body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return SessionInfo{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		return SessionInfo{}, err
	}
	defer res.Body.Close()
	// read body
	responseBody, err := io.ReadAll(res.Body)
	fmt.Println(string(responseBody))
	if err != nil {
		return SessionInfo{}, err
	}

	info := SessionInfo{}
	err = xml.Unmarshal(responseBody, &info)
	if err != nil {
		return SessionInfo{}, err
	}
	return info, nil
}

func login() (string, error) {
	info, err := sendLoginRequest("GET", nil)
	if err != nil {
		return "", err
	}

	fmt.Println(info)
	response := challenge(info.Challenge, Password)
	data := url.Values{}
	data.Set("response", response)
	data.Set("username", Username)
	info, err = sendLoginRequest("POST", strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	fmt.Println(info)
	return info.SID, nil
}

func logout(sid string) error {
	data := url.Values{}
	data.Set("logout", sid)

	info, err := sendLoginRequest("POST", strings.NewReader(data.Encode()))
	fmt.Println(info)
	return err
}

func check(sid string) error {
	data := url.Values{}
	data.Set("sid", sid)

	info, err := sendLoginRequest("POST", strings.NewReader(data.Encode()))
	fmt.Println(info)
	return err
}

func main() {
	sid, err := login()
	if err != nil {
		panic(err)
	}
	fmt.Println(sid)
	check(sid)
	logout(sid)
}
