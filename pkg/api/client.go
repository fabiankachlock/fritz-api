package api

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

	"github.com/fabiankachlock/fritz-api/pkg/response"
	"golang.org/x/crypto/pbkdf2"
)

type Client struct {
	boxUrl string
	sid    string
}

func NewClient(boxUrl string) *Client {
	return &Client{boxUrl: boxUrl}
}

func (c *Client) Login(username string, password string) error {
	info, err := c.sendLoginRequest("GET", nil)
	if err != nil {
		return err
	}

	data := url.Values{}
	response := c.solveChallenge(info.Challenge, password)
	data.Set("response", response)
	data.Set("username", username)

	info, err = c.sendLoginRequest("POST", data)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Logout() error {
	data := url.Values{}
	data.Set("logout", c.sid)

	_, err := c.sendLoginRequest("POST", data)
	return err
}

func (c *Client) CheckLogin() error {
	data := url.Values{}
	data.Set("sid", c.sid)

	_, err := c.sendLoginRequest("POST", data)
	return err
}

func (c *Client) GetSession() (response.SessionInfo, error) {
	data := url.Values{}
	data.Set("sid", c.sid)

	return c.sendLoginRequest("POST", data)
}

func (c Client) sendLoginRequest(method string, body url.Values) (response.SessionInfo, error) {
	loginUrl := fmt.Sprintf("%s/login_sid.lua?version=2", c.boxUrl)
	req, err := http.NewRequest(method, loginUrl, strings.NewReader(body.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return response.SessionInfo{}, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return response.SessionInfo{}, err
	}
	responseBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return response.SessionInfo{}, err
	}

	info := response.SessionInfo{}
	err = xml.Unmarshal(responseBody, &info)
	if err != nil {
		return response.SessionInfo{}, err
	}
	return info, nil
}

func (c Client) solveChallenge(challenge string, pw string) string {
	parts := strings.Split(challenge, "$")
	iter1, _ := strconv.Atoi(parts[1])
	iter2, _ := strconv.Atoi(parts[3])
	salt1, _ := hex.DecodeString(parts[2])
	salt2, _ := hex.DecodeString(parts[4])
	hash1 := pbkdf2.Key([]byte(pw), salt1, iter1, 32, sha256.New)
	hash2 := pbkdf2.Key(hash1, salt2, iter2, 32, sha256.New)
	return parts[4] + "$" + hex.EncodeToString(hash2)
}
