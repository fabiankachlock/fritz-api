package request

import "github.com/fabiankachlock/fritz-api/pkg/request/transformer"

type DataRequest struct {
	Parameters          map[string]string
	ResponseTransformer func(content []byte) ([]byte, error)
}

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
)
