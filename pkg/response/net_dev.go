package response

// NetDev is the response of requests send by the network devices [Site]
//
// [Site]: http://fritz.box/#netDev
type NetDev struct {
	IPClient    bool           `json:"ipclient"`    // IPClient [unidentified]
	IsRepeater  bool           `json:"isrepeater"`  // IsRepeater [unidentified]
	IsPowerline bool           `json:"ispowerline"` // IsPowerline [unidentified]
	Initial     bool           `json:"initial"`     // Initial [unidentified]
	NexusClient bool           `json:"nexusclient"` // NexusClient [unidentified]
	BackToPage  string         `json:"backToPage"`  // BackToPage [unidentified]
	BridgeMode  BridgeModeType `json:"bridgeMode"`  // BridgeMode [unidentified]

	FritzBoxOther []NetworkDevice `json:"fbox_other"` // FritzBoxOther [unidentified]
	FritzBox      []NetworkDevice `json:"fbox"`       // FritzBox [unidentified]

	Active  []NetworkDevice `json:"active"`  // Active [unidentified]
	Passive []NetworkDevice `json:"passive"` // Passive [unidentified]

	FritzBoxTitle     string `json:"fbox_title"`     // FritzBoxTitle [unidentified]
	TitleDeviceAdd    string `json:"titledeviceadd"` // TitleDeviceAdd [unidentified]
	TitleDeviceDelete string `json:"titledevicedel"` // TitleDeviceDelete [unidentified]
}

// BridgeModeType [unidentified]
type BridgeModeType string

const (
	BridgeModeTypeWlan BridgeModeType = "wlan_bridge"
)

type NetworkDevice struct {
	UID             string             `json:"UID"`               // UID is the unique identifier of the device
	Mac             string             `json:"mac"`               // Mac is the mac address of the device
	OwnClientDevice bool               `json:"own_client_device"` // OwnClientDevice is whether this device is the currently interacting client
	Type            NetworkDeviceClass `json:"type"`              // Type [unidentified] seems to be the type of the connection
	Classes         NetworkDeviceClass `json:"classes"`           // Classes [unidentified] seems to be always the same as type
	Port            string             `json:"port"`              // Port on which the device is connected to the fritz box (eg. WLAN or LAN 2 mit 1 Gbit/s)
	Name            string             `json:"name"`              // Name of the device
	Model           ModelType          `json:"model"`             // Model [unidentified] seems to be the group of the device
	BoxType         BoxType            `json:"boxType"`           // BoxType [unidentified] seems to be the type of the fritz box
	URL             string             `json:"url"`               // URL is the URL to the device in the web ui

	Connections []InternetConnection `json:"connections"` // Connections is a list of all connections the device has to the internet (only fbox devices)

	Options    NetworkDeviceOptions    `json:"options"`                         // Options are additional options of the device displayed in the ui
	Parent     NetworkDeviceParent     `json:"parent"`                          // Parent [unidentified]
	Properties []NetworkDeviceProperty `json:"properties"`                      // Properties are additional properties of the device displayed in the ui
	State      NetworkDeviceStateInfo  `json:"state" transform:"extendedEmpty"` // State is the connection state displayed in the ui
	IPv4       IPv4Info                `json:"ipv4"`                            // IPv4 is the ipv4 address of the device
}

// NetworkDeviceClass classifies the devices by connection type
type NetworkDeviceClass string

const (
	NetWorkDeviceClassWlan     NetworkDeviceClass = "wlan"
	NetWorkDeviceClassEthernet NetworkDeviceClass = "ethernet"
	NetWorkDeviceClassUnknown  NetworkDeviceClass = "unknown"
)

// ModelType [unidentified] seems to be the group of the device
type ModelType string

const (
	ModelTypeFritzBox ModelType = "fbox"
	ModelTypeActive   ModelType = "active"
	ModelTypePassive  ModelType = "passive"
)

// NetworkDeviceStateClass classifies the connection state of the device
type NetworkDeviceStateClass string

const (
	NetworkDeviceStateClassLedGreen    NetworkDeviceStateClass = "led_green"
	NetworkDeviceStateClassGlobeOnline NetworkDeviceStateClass = "globe_online"
)

// Options are additional options of the device displayed in the ui
type NetworkDeviceOptions struct {
	Guest      bool `json:"guest"`
	Editable   bool `json:"editable"`
	Deleteable bool `json:"deleteable"`
	Disable    bool `json:"disable"`
}

// Parent [unidentified]
type NetworkDeviceParent struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Properties are additional properties of the device displayed in the ui
type NetworkDeviceProperty struct {
	Txt     string `json:"txt"`
	OnClick string `json:"onclick"`
	Icon    string `json:"icon"`
	Link    string `json:"link"`
}

// State is the connection state displayed in the ui
type NetworkDeviceStateInfo struct {
	Class NetworkDeviceStateClass `json:"class"`
}

// IPv4 is the ipv4 address of the device
type IPv4Info struct {
	Node        string `json:"_node"`    // Node [unidentified]
	AddressType string `json:"addrtype"` // AddressType is the type of the address (eg. ipv4)
	Dhcp        string `json:"dhcp"`     // DHCP is whether the device uses DHCP (values: "1" or "0")
	LastUsed    string `json:"lastused"` // LastUsed is the last time the device was used
	IP          string `json:"ip"`       // IP is the ipv4 address of the device
}
