package response

type Energy struct {
	Intro struct {
		Text []string `json:"text"`
	} `json:"intro"`
	Drain []EnergyDrainInfo `json:"drain"`
}

type EnergyDrainInfo struct {
	CurrentPercentage     int             `json:"actPerc"`                      // CurrentPercentage is the current percentage of energy usage
	AccumulatedPercentage int             `json:"cumPerc"`                      // AccumulatedPercentage is the accumulated percentage of the last 24 hours
	Name                  string          `json:"name"`                         // Name is the name of the metric
	Statuses              []string        `json:"statuses" transform:"toSlice"` // Statuses is the status of the metric
	Lan                   []EnergyLanInfo `json:"lan"`                          // Lan is the status of the lan ports (usually only in the last element)
}

type EnergyLanInfo struct {
	Name  string       `json:"name"`
	Class LanInfoClass `json:"class"`
}

type LanInfoClass string

const (
	LanInfoClassGreen LanInfoClass = "green"
)
