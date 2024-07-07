package response

// HomeNet is the response of requests send by the mesh [Site]
//
// [Site]: http://fritz.box/#homeNet
type HomeNet struct {
	Searching bool `json:"searching"` // Searching [unidentified]
	IPClient  bool `json:"ipclient"`  // IPClient [unidentified]
	// Updating [unidentified]
	// NOTE: this may be connected to a request field 'updating'
	// it may be in realtion to probably partial ui updates but the
	// responses don't seem to be reduced in size
	Updating      string        `json:"updating"`
	Devices       []MeshDevice  `json:"devices"`     // Devices is a list of all devices currently connected to the (mesh)-network
	Topology      Topology      `json:"topology"`    // Topology is a hierarchical representation of the currently connected devices in the network
	FirmwareCheck FirmwareCheck `json:"fwcheck"`     // FirmwareCheck [unidentified] seems to be status of the firmware check
	NexusClient   bool          `json:"nexusclient"` // NexusClient [unidentified]
}

// FirmwareCheck is the status of the firmware check
type FirmwareCheck struct {
	NotDone bool `json:"notdone"`
	NoCheck bool `json:"nocheck"`
	Auto    bool `json:"auto"`
	Started bool `json:"started"`
}

// Topology is a hierarchical representation of the currently connected devices in the network
type Topology struct {
	RootDeviceId string                `json:"rootuid"`
	Devices      map[string]MeshDevice `json:"devices"`
}

// MeshDevice is a device connected to the network
type MeshDevice struct {
	UID             string         `json:"UID"`               // UID is the unique identifier of the device in the mesh topology
	OwnEntry        bool           `json:"ownentry"`          // OwnEntry is whether the device is itself
	OwnClientDevice bool           `json:"own_client_device"` // OwnClientDevice is whether this device is the currently interacting client
	Master          bool           `json:"master"`            // Master is whether the device is the master device
	Dist            int            `json:"dist"`              // Dist [unidentified] this seems to be the distance (in nodes) to the root device in the mesh
	Parent          string         `json:"parent"`            // Parent the parent device in the topology view
	Children        []string       `json:"children"`          // Children is a list of all devices connected to this device
	Category        CategoryType   `json:"category"`          // Category the category of the device in the topology (ownentry for the root device)
	ConnectionType  ConnectionType `json:"conn"`              // ConnectionType the type of connection (any for the root device)
	Switch          bool           `json:"switch"`            // Switch [unidentified] might be whether the devices acts like a switch
	DevType         DevType        `json:"devtype"`           // DevType the type of the device
	IPInfo          string         `json:"ipinfo"`            // IPInfo [unidentified] seems to be the ip address of the device
	Gateway         bool           `json:"gateway"`           // Gateway [unidentified]
	BoxType         BoxType        `json:"boxType"`           // BoxType [unidentified] seems to be the type of the fritz box

	NameInfo    NameInfo            `json:"nameinfo"`                              // NameInfo is the name information of the device
	VersionInfo VersionInfo         `json:"versioninfo" transform:"extendedEmpty"` // VersionInfo is the version information of the device
	StateInfo   MeshDeviceStateInfo `json:"stateinfo"`                             // StateInfo is the state information of the device
	UpdateInfo  UpdateInfo          `json:"updateinfo"`                            // UpdateInfo is the software update information of the device (none for unknown/"normal" devices)
	PhoneInfo   PhoneInfo           `json:"phone"`                                 // PhoneInfo is the phone information of the device (only for the root device)
	WlanInfo    []WlanInfo          `json:"wlaninfo"`                              // WlanInfo is the information about wlan networks exposed by the device
	Detailinfo  Detailinfo          `json:"detailinfo"`                            // Detailinfo is information about details of the device (in the ui)
	ConnInfo    ConnectionInfo      `json:"conninfo" transform:"extendedEmpty"`    // ConnectionInfo is the information about the connection to the router

	Connections []InternetConnection `json:"connections"` // Connections is a list of all connections the device has to the internet (only for the root device)
}

// CategoryType is the type of category the device is in the topology
type CategoryType string

const (
	CategoryOwn  CategoryType = "ownentry"
	CategoryLan  CategoryType = "lan"
	CategoryWlan CategoryType = "wlan"
)

// ConnectionType is the type of connection the device has to the network
type ConnectionType string

const (
	ConnectionTypeAny  ConnectionType = "any"
	ConnectionTypeLan  ConnectionType = "lan"
	ConnectionTypeWlan ConnectionType = "wlan"
)

type BoxType string

const (
	BoxTypeDSL BoxType = "dsl"
)

type DevType string

const (
	DevTypeFritzBox BoxType = "fritzbox"
)

// VersionInfo information about the version of the device
type VersionInfo struct {
	Version string `json:"version"` // Version is the version of the device
	Fos     bool   `json:"fos"`     // Fos [unidentified]
}

// NameInfo is information about the name of the device
type NameInfo struct {
	Name    string `json:"name"`    // Name is the name of the device
	Product string `json:"product"` // Product is the product name of the device (only when the device is a FRITZ!-device)
	URL     string `json:"url"`     // URL is an optional url to the device (only when the device has a web interface)
}

// MeshDeviceStateInfo is information about the state of the device
type MeshDeviceStateInfo struct {
	Active bool `json:"active"` // Active whether the device is currently active

	// the following section only applies "normal" client

	GuestOwe        bool `json:"guest_owe"`       // GuestOwe [unidentified]
	Meshable        bool `json:"meshable"`        // Meshable [unidentified]
	Guest           bool `json:"guest"`           // Guest [unidentified] may be whether the device is a guest
	Online          bool `json:"online"`          // Online whether the device is currently online
	Blocked         bool `json:"blocked"`         // Blocked [unidentified]
	Realtime        bool `json:"realtime"`        // Realtime [unidentified]
	NotAllowed      bool `json:"notallowed"`      // NotAllowed [unidentified]
	InternetBlocked bool `json:"internetBlocked"` // InternetBlocked [unidentified]

	// the following section only applies to the root(?) device

	NexusTrust bool `json:"nexustrust"` // NexusTrust [unidentified] {only when the device is a FRITZ!-device}
}

// UpdateInfo is the software update information of the device
type UpdateInfo struct {
	State UpdateInfoState `json:"state"` // State is the state of the software update info
}

type UpdateInfoState string

const (
	UpdateInfoStateCurrent UpdateInfoState = "current"
	UpdateInfoStateNone    UpdateInfoState = "none"
)

// PhoneInfo is information about available phone numbers of the device
type PhoneInfo struct {
	NumberCount int `json:"numberCount"` // NumberCount is the number of available phone numbers
	ActiveCount int `json:"activeCount"` // ActiveCount is the number of active phone numbers
}

// WlanInfo is the information about wlan networks exposed by the device
type WlanInfo struct {
	Text       string `json:"text"`       // Text is the name (SSID) of the wlan network
	Title      string `json:"title"`      // Title is the title of the wlan network (ex. WLAN-Funknetz or Gastfunknetz)
	ShortTitle string `json:"shorttitle"` // ShortTitle is a shorter version of the title
}

// DetailInfo is information about details of the device (in the ui)
type Detailinfo struct {
	// the following section only applies to the root(?) device

	Wlan24      bool `json:"wlan24"`      // Wlan24 whether the 2.4GHz wlan is enabled
	Wlan5       bool `json:"wlan5"`       // Wlan5 whether the 5GHz wlan is enabled
	GuestAccess bool `json:"guestaccess"` // GuestAccess whether a guest network is enabled

	// the following section only applies "normal" client

	Edit struct {
		// Pid [unidentified] seems to be the page id of the edit page
		Pid string `json:"pid"`
		// Params [unidentified] seem to be url parameters for navigation actions
		// known params are:
		//   - dev -> device.UID
		//   - back_to_page -> homeNet (current page id)
		Params map[string]string `json:"params"`
	} `json:"edit"`
	PortRelease bool `json:"portrelease"`
}

// ConnectionInfo is the information about the connection to the router
type ConnectionInfo struct {
	Kind  ConnectionType `json:"kind"`  // Kind is the type of connection
	Speed string         `json:"speed"` // Speed is the speed of the connection with unit (ex. 100 MBit/s or 1 GBit/s)
	Desc  string         `json:"desc"`  // Desc is a description of the connection (ex. LAN 1 or 2.4GHz, 5GHz)

	UsedBands int `json:"usedbands"` // UsedBands is the number of wlan bands used by the device
	// BandInfo is the information about the wlan band used
	BandInfo []struct {
		Band    int    `json:"band"`     // Band is the band number
		SpeedTx int    `json:"speed_tx"` // SpeedTx is the speed of the connection in transmit direction in mbit/s
		SpeedRx int    `json:"speed_rx"` // SpeedRx is the speed of the connection in receive direction in mbit/s
		Speed   string `json:"speed"`    // Speed is the speed of the connection with unit (ex. 100 MBit/s or 1 GBit/s)
		Desc    string `json:"desc"`     // Desc is a description of the connection (ex. 2.4GHz, 5GHz)
	} `json:"bandinfo"`
}

// InternetConnection is the connection information of the device
type InternetConnection struct {
	Name             string                   `json:"name"`
	Connected        bool                     `json:"connected"`
	Active           bool                     `json:"active"`
	DslDiagnosis     bool                     `json:"dsl_diagnosis"`
	DirectConnection bool                     `json:"direct_connection"`
	ReadyForFallback bool                     `json:"ready_for_fallback"`
	ShapedRate       bool                     `json:"shapedrate"`
	SpeedManual      bool                     `json:"speed_manual"`
	Downstream       int                      `json:"downstream"`        // in kbit/s
	MediumDownstream int                      `json:"medium_downstream"` // in kbit/s
	Upstream         int                      `json:"upstream"`          // in kbit/s
	MediumUpstream   int                      `json:"medium_upstream"`   // in kbit/s
	Role             string                   `json:"role"`
	Provider         string                   `json:"provider"`
	State            InternetConnectionState  `json:"state"`
	Type             InternetConnectionType   `json:"type"`
	Medium           InternetConnectionMedium `json:"medium"`

	IPv4 struct {
		Connected bool      `json:"connected"`
		IP        string    `json:"ip"`
		Since     int       `json:"since"`
		DNS       []DNSInfo `json:"dns"`
		DSLite    bool      `json:"dslite"`
		Ipv6AFTR  string    `json:"ipv6_aftr"`
	} `json:"ipv4"`
	Ipv6 struct {
		Connected  bool   `json:"connected"`
		IP         string `json:"ip"`
		Since      int    `json:"since"`
		IPLifetime struct {
			Valid     int `json:"valid"`
			Preferred int `json:"preferred"`
		} `json:"ip_lifetime"`
		Prefix         string `json:"prefix"`
		PrefixLifetime struct {
			Valid     int `json:"valid"`
			Preferred int `json:"preferred"`
		} `json:"prefix_lifetime"`
		DNS []DNSInfo `json:"dns"`
	} `json:"ipv6"`
}

type InternetConnectionState string

const (
	InternetConnectionStateConnected InternetConnectionState = "connected"
)

type InternetConnectionType string

const (
	InternetConnectionTypeDSL InternetConnectionType = "dsl"
)

type InternetConnectionMedium string

const (
	InternetConnectionMediumDSL InternetConnectionMedium = "dsl"
)

type DNSInfo struct {
	Edns0Mode string `json:"edns0_mode"`
	Type      string `json:"type"`
	IP        string `json:"ip"`
	Purpose   string `json:"purpose"`
}
