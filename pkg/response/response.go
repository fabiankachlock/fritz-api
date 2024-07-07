package response

import "encoding/json"

type DataResponse[T any] struct {
	PageId string `json:"pid"`
	// TODO: figure out what 'hide' does and how its defined
	// It seems to be a map of true values configuring a set of ui options that should be hidden
	// Hide           map[string]bool `json:"hide"`
	TimeTillLogout int `json:"timeTillLogout,string"`
	// TODO: figure out what 'time' does and how its defined
	// Until now it only saw empty arrays for this field
	// Time           []any           `json:"time"`
	Data T      `json:"data"`
	Sid  string `json:"sid"`
}

func UnmarshalAs[T any](response []byte) (DataResponse[T], error) {
	var dataResponse DataResponse[T]
	err := json.Unmarshal(response, &dataResponse)
	if err != nil {
		return dataResponse, err
	}
	return dataResponse, nil
}
