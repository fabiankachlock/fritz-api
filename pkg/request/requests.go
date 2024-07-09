package request

import "github.com/fabiankachlock/fritz-api/pkg/request/transformer"

type DataRequest struct {
	Parameters          map[string]string
	ResponseTransformer func(content []byte) ([]byte, error)
}

const (
	RequestParamFilter = "filter"
)

var (
	MeshRequest = DataRequest{
		Parameters: map[string]string{
			"page": "homeNet",
		},
		ResponseTransformer: transformer.Noop,
	}

	NetworkDevicesRequest = DataRequest{
		Parameters: map[string]string{
			"page": "netDev",
		},
		ResponseTransformer: transformer.Noop,
	}

	NetworkUsageRequest = DataRequest{
		Parameters: map[string]string{
			"page": "netCnt",
		},
		ResponseTransformer: transformer.NetCntToJson,
	}

	EnergyUsageRequest = DataRequest{
		Parameters: map[string]string{
			"page": "energy",
		},
		ResponseTransformer: transformer.Noop,
	}

	SystemLogsRequest = DataRequest{
		Parameters: map[string]string{
			"page":             "log",
			RequestParamFilter: "all",
		},
		ResponseTransformer: transformer.Noop,
	}
)

func WithParams(request DataRequest, params map[string]string) DataRequest {
	for key, value := range params {
		request.Parameters[key] = value
	}
	return request
}
