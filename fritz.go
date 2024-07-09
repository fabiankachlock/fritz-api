package fritz

import (
	"github.com/fabiankachlock/fritz-api/pkg/api"
	"github.com/fabiankachlock/fritz-api/pkg/request"
	"github.com/fabiankachlock/fritz-api/pkg/response"
)

type FritzBoxClient struct {
	client *api.Client
}

// NewClient creates a new FritzBoxClient
func NewClient(boxUrl string) FritzBoxClient {
	return FritzBoxClient{client: api.NewClient(boxUrl)}
}

// Login logs in to the FritzBox
func (f *FritzBoxClient) Login(username string, password string) error {
	return f.client.Login(username, password)
}

// Logout logs out of the FritzBox
func (f *FritzBoxClient) Logout() error {
	return f.client.Logout()
}

// CheckLogin checks if the client is logged in
func (f *FritzBoxClient) CheckLogin() error {
	return f.client.CheckLogin()
}

// GetSession returns the current session
func (f *FritzBoxClient) GetSession() (response.SessionInfo, error) {
	return f.client.GetSession()
}

// GetData returns the response body for a custom data request
func (f *FritzBoxClient) GetData(request request.DataRequest) ([]byte, error) {
	return f.client.RequestData(request)
}

// GetDataJson returns the response body as marshalled json for a custom data request as a map
func (f *FritzBoxClient) GetDataJson(request request.DataRequest) (map[string]interface{}, error) {
	return PerformRequestRaw(f, request)
}

// GetHomeNet returns the data for the home network page
func (f *FritzBoxClient) GetHomeNet() (response.DataResponse[response.HomeNet], error) {
	return PerformRequest[response.HomeNet](f, request.MeshRequest)
}

func (f *FritzBoxClient) GetNetworkDevices() (response.DataResponse[response.NetDev], error) {
	return PerformRequest[response.NetDev](f, request.NetworkDevicesRequest)
}

func (f *FritzBoxClient) GetNetworkUsage() (response.NetCnt, error) {
	return PerformRequestCustom[response.NetCnt](f, request.NetworkUsageRequest)
}

func (f *FritzBoxClient) GetEnergyUsage() (response.Energy, error) {
	return PerformRequestCustom[response.Energy](f, request.EnergyUsageRequest)
}

// PerformRequest is a wrapper that performs a request and unmarshals the response as data response
func PerformRequest[T any](f *FritzBoxClient, req request.DataRequest) (response.DataResponse[T], error) {
	body, err := f.client.RequestData(req)
	if err != nil {
		return response.DataResponse[T]{}, err
	}

	json, err := response.UnmarshalAs[T](body)
	if err != nil {
		return response.DataResponse[T]{}, err
	}
	return json, nil
}

// PerformRequestCustom is a wrapper that performs a request and unmarshals the response
func PerformRequestCustom[T any](f *FritzBoxClient, req request.DataRequest) (T, error) {
	var empty T
	body, err := f.client.RequestData(req)
	if err != nil {
		return empty, err
	}

	json, err := response.UnmarshalCustomAs[T](body)
	if err != nil {
		return empty, err
	}
	return json, nil
}

// PerformRequestRaw is a wrapper that performs a request and unmarshals the response as raw json
func PerformRequestRaw(f *FritzBoxClient, req request.DataRequest) (map[string]interface{}, error) {
	body, err := f.client.RequestData(req)
	if err != nil {
		return map[string]interface{}{}, err
	}

	rawJson, err := response.UnmarshalRaw(body)
	if err != nil {
		return map[string]interface{}{}, err
	}
	return rawJson, nil
}
