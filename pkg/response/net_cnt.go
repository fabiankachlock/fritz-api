package response

type NetCnt struct {
	LastMonth UsageMetric `json:"LastMonth"`
	ThisWeek  UsageMetric `json:"ThisWeek"`
	Today     UsageMetric `json:"Today"`
	Yesterday UsageMetric `json:"Yesterday"`
	ThisMonth UsageMetric `json:"ThisMonth"`
}

type UsageMetric struct {
	BytesSentHigh     int64 `json:"BytesSentHigh" transform:"stringToInt"`
	BytesSentLow      int64 `json:"BytesSentLow" transform:"stringToInt"`
	BytesReceivedHigh int64 `json:"BytesReceivedHigh" transform:"stringToInt"`
	BytesReceivedLow  int64 `json:"BytesReceivedLow" transform:"stringToInt"`
}
