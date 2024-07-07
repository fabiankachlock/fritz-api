package response

import (
	"encoding/json"

	"github.com/fabiankachlock/fritz-api/pkg/helper/transform"
)

type DataResponse[T any] struct {
	PageId string `json:"pid"`
	// TODO: figure out what 'hide' does and how its defined
	// It seems to be a map of true values configuring a set of ui options that should be hidden
	Hide           map[string]bool `json:"hide"`
	TimeTillLogout string          `json:"timeTillLogout"`
	// TODO: figure out what 'time' does and how its defined
	// Until now it only saw empty arrays for this field
	Time []any  `json:"time"`
	Data T      `json:"data"`
	Sid  string `json:"sid"`
}

func UnmarshalRaw(response []byte) (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := json.Unmarshal(response, &data)
	if err != nil {
		return map[string]interface{}{}, err
	}

	return data, nil
}

func UnmarshalAs[T any](response []byte) (DataResponse[T], error) {
	var dataResponse DataResponse[T]
	data := map[string]interface{}{}
	err := json.Unmarshal(response, &data)
	if err != nil {
		return dataResponse, err
	}

	return transform.MapToStruct(data, dataResponse), nil
}
