package request

type DataRequest struct {
	Parameters map[string]string
}

var (
	MeshRequest = DataRequest{
		Parameters: map[string]string{
			"page": "homeNet",
		},
	}

	NetworkDevicesRequest = DataRequest{
		Parameters: map[string]string{
			"page": "netDev",
		},
	}
)
