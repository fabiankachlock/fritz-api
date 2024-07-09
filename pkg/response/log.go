package response

type Log struct {
	// Show seems to be the available filter options
	Show struct {
		All     bool `json:"all"`
		System  bool `json:"sys"`
		Network bool `json:"net"`
		Phone   bool `json:"fon"`
		USB     bool `json:"usb"`
		WLAN    struct {
			WPA2Support bool `json:"has_wpa2_support"`
			WPA3Support bool `json:"has_wpa3_support"`
		} `json:"wlan"`
	} `json:"show"`
	Logs              []LogEntry `json:"log"`               // Logs the system logs
	Filter            LogFilter  `json:"filter"`            // Filter the currently applied log filter
	WLANGuestPushMail bool       `json:"wlanGuestPushmail"` // WLANGuestPushMail [unidentified]
	WLAN              bool       `json:"wlan"`              // WLAN [unidentified]
}

type LogEntry struct {
	ID      int       `json:"id"`     // ID seem to be the type of the log entry (the help link will show help according to the id)
	Time    string    `json:"time"`   // Time is the time of the log entry
	Group   LogFilter `json:"group"`  // Group is the group of the log entry
	Message string    `json:"msg"`    // Message is the message of the log entry
	Date    string    `json:"date"`   // Date is the date of the log entry
	NoHelp  int       `json:"noHelp"` // NoHelp seems to be a flag if the help link should be shown
}

type LogFilter string

const (
	LogFilterAll  LogFilter = "all"
	LogFilterSys  LogFilter = "sys"
	LogFilterNet  LogFilter = "net"
	LogFilterFon  LogFilter = "fon"
	LogFilterUSB  LogFilter = "usb"
	LogFilterWLAN LogFilter = "wlan"
)
