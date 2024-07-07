package response

// HomeNet is the response of requests send by the homeNet [Site]
//
// [Site]: http://fritz.box/#homeNet
type HomeNet struct {
	// Searching [unidentified]
	Searching bool `json:"searching"`
	// IPClient [unidentified]
	IPClient bool `json:"ipclient"`
	// Updating [unidentified]
	// NOTE: this may be connected to a request field 'updating'
	// it may be in realtion to probably partial ui updates but the
	// responses don't seem to be reduced in size
	Updating string `json:"updating"`
	// Devices is a list of all devices currently connected to the (mesh)-network
	Devices []HomeNetDevice `json:"devices"`
	// Topology is a hierarchical representation of the currently connected devices in the network
	Topology Topology `json:"topology"`
	// FirmwareCheck [unidentified] seems to be status of the firmware check
	FirmwareCheck FirmwareCheck `json:"fwcheck"`
	// NexusClient [unidentified]
	NexusClient bool `json:"nexusclient"`
}

type FirmwareCheck struct {
	NotDone bool `json:"notdone"`
	NoCheck bool `json:"nocheck"`
	Auto    bool `json:"auto"`
	Started bool `json:"started"`
}

type Topology struct {
	RootDeviceId string                    `json:"rootuid"`
	Devices      map[string]TopologyDevice `json:"devices"`
}

type TopologyDevice struct {
	OwnClientDevice bool   `json:"own_client_device"` // [unidentified]
	Dist            int    `json:"dist"`              // [unidentified] this seems toi be the distance (in node) toi the root device
	Parent          string `json:"parent"`            // the parent device in the topology view
	VersionInfo     []struct {
		Version string `json:"version"`
		Fos     bool   `json:"fos"`
	} `json:"versioninfo"`
	UID       string   `json:"UID"`
	Category  string   `json:"category"`
	Switch    bool     `json:"switch"`
	Children  []string `json:"children"`
	Devtype   string   `json:"devtype"`
	Ownentry  bool     `json:"ownentry"`
	Stateinfo struct {
		GuestOwe        bool `json:"guest_owe"`
		Active          bool `json:"active"`
		Meshable        bool `json:"meshable"`
		Guest           bool `json:"guest"`
		Online          bool `json:"online"`
		Blocked         bool `json:"blocked"`
		Realtime        bool `json:"realtime"`
		Notallowed      bool `json:"notallowed"`
		InternetBlocked bool `json:"internetBlocked"`
		Nexustrust      bool `json:"nexustrust"`
	} `json:"stateinfo"`
	Connections []Connection `json:"connections"`
	Conn        string       `json:"conn"`
	Master      bool         `json:"master"`
	Ipinfo      string       `json:"ipinfo"`
	Updateinfo  struct {
		State string `json:"state"`
	} `json:"updateinfo"`
	Nameinfo struct {
		Name    string `json:"name"`
		Product string `json:"product"`
		URL     string `json:"url"`
	} `json:"nameinfo"`
	Gateway bool   `json:"gateway"`
	BoxType string `json:"boxType"`
	Phone   struct {
		NumberCount int `json:"numberCount"`
		ActiveCount int `json:"activeCount"`
	} `json:"phone"`
	Wlaninfo []struct {
		Text       string `json:"text"`
		Title      string `json:"title"`
		Shorttitle string `json:"shorttitle,omitempty"`
	} `json:"wlaninfo"`
	Detailinfo struct {
		Wlan24      bool `json:"wlan24"`
		Wlan5       bool `json:"wlan5"`
		Guestaccess bool `json:"guestaccess"`
		Edit        struct {
			Pid    string `json:"pid"`
			Params struct {
				Dev        string `json:"dev"`
				BackToPage string `json:"back_to_page"`
			} `json:"params"`
		} `json:"edit"`
		Portrelease bool `json:"portrelease"`
	} `json:"detailinfo"`
	Conninfo struct {
		Kind     string `json:"kind"`
		Speed    string `json:"speed"`
		Bandinfo []struct {
			Band    int    `json:"band"`
			SpeedTx int    `json:"speed_tx"`
			SpeedRx int    `json:"speed_rx"`
			Speed   string `json:"speed"`
			Desc    string `json:"desc"`
		} `json:"bandinfo"`
		Usedbands int    `json:"usedbands"`
		Desc      string `json:"desc"`
	} `json:"conninfo"`
}

type Connection struct {
	DslDiagnosis   bool   `json:"dsl_diagnosis"`
	MediumUpstream int    `json:"medium_upstream"`
	Downstream     int    `json:"downstream"`
	Role           string `json:"role"`
	Provider       string `json:"provider"`
	Ipv4           struct {
		Ipv6Aftr  string `json:"ipv6_aftr"`
		Connected bool   `json:"connected"`
		DNS       []any  `json:"dns"`
		Dslite    bool   `json:"dslite"`
		IP        string `json:"ip"`
		Since     int    `json:"since"`
	} `json:"ipv4"`
	Connected        bool   `json:"connected"`
	Shapedrate       bool   `json:"shapedrate"`
	DirectConnection bool   `json:"direct_connection"`
	ReadyForFallback bool   `json:"ready_for_fallback"`
	MediumDownstream int    `json:"medium_downstream"`
	State            string `json:"state"`
	Upstream         int    `json:"upstream"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	Active           bool   `json:"active"`
	Ipv6             struct {
		IPLifetime struct {
			Valid     int `json:"valid"`
			Preferred int `json:"preferred"`
		} `json:"ip_lifetime"`
		Connected bool `json:"connected"`
		DNS       []struct {
			Edns0Mode string `json:"edns0_mode"`
			Type      string `json:"type"`
			IP        string `json:"ip"`
			Purpose   string `json:"purpose,omitempty"`
		} `json:"dns"`
		IP             string `json:"ip"`
		Prefix         string `json:"prefix"`
		PrefixLifetime struct {
			Valid     int `json:"valid"`
			Preferred int `json:"preferred"`
		} `json:"prefix_lifetime"`
		Since int `json:"since"`
	} `json:"ipv6"`
	SpeedManual bool   `json:"speed_manual"`
	Medium      string `json:"medium"`
}

type HomeNetDevice struct {
	OwnClientDevice bool   `json:"own_client_device"`
	Dist            int    `json:"dist"`
	Parent          string `json:"parent"`
	VersionInfo     struct {
		Version string `json:"version"`
		Fos     bool   `json:"fos"`
	} `json:"versioninfo" transform:"extendedEmpty"`
	UID      string   `json:"UID"`
	Switch   bool     `json:"switch"`
	Children []string `json:"children"`
	Wlaninfo []struct {
		Text       string `json:"text"`
		Title      string `json:"title"`
		Shorttitle string `json:"shorttitle"`
	} `json:"wlaninfo"`
	Connections []Connection `json:"connections"`
	Conninfo    []any        `json:"conninfo"`
	Ipinfo      string       `json:"ipinfo"`
	Devtype     string       `json:"devtype"`
	Ownentry    bool         `json:"ownentry"`
	Phone       struct {
		NumberCount int `json:"numberCount"`
		ActiveCount int `json:"activeCount"`
	} `json:"phone"`
	Conn     string `json:"conn"`
	Master   bool   `json:"master"`
	BoxType  string `json:"boxType"`
	Gateway  bool   `json:"gateway"`
	Nameinfo struct {
		Name    string `json:"name"`
		Product string `json:"product"`
		URL     string `json:"url"`
	} `json:"nameinfo"`
	Updateinfo struct {
		State string `json:"state"`
	} `json:"updateinfo"`
	Category   string `json:"category"`
	DetailInfo struct {
		Wlan24      bool `json:"wlan24"`
		Wlan5       bool `json:"wlan5"`
		Guestaccess bool `json:"guestaccess"`
		Edit        struct {
			Pid    string `json:"pid"`
			Params struct {
				Dev        string `json:"dev"`
				BackToPage string `json:"back_to_page"`
			} `json:"params"`
		} `json:"edit"`
		PortRelease bool `json:"portrelease"`
	} `json:"detailinfo"`
	StateInfo struct {
		Nexustrust      bool `json:"nexustrust"`
		GuestOwe        bool `json:"guest_owe"`
		Active          bool `json:"active"`
		Meshable        bool `json:"meshable"`
		Guest           bool `json:"guest"`
		Online          bool `json:"online"`
		Blocked         bool `json:"blocked"`
		Realtime        bool `json:"realtime"`
		Notallowed      bool `json:"notallowed"`
		InternetBlocked bool `json:"internetBlocked"`
	} `json:"stateinfo"`
}
