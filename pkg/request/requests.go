package request

type DataRequest struct {
	Parameters map[string]string
}

var (
	HomeNetRequest = DataRequest{
		Parameters: map[string]string{
			"page": "homeNet",
		},
	}
)
