package bsn

type Switch struct {
	Connected               bool        `json:"connected"`
	ConnectedSince          string      `json:"connected-since"`
	DPID                    string      `json:"dpid"`
	FabricConnectionState   string      `json:"fabric-connection-state"`
	FabricLastSeenTime      string      `json:"fabric-last-seen-time"`
	FabricRole              string      `json:"fabric-role"`
	HandshakeState          string      `json:"handshake-state"`
	InetAddress             InetAddress `json:"inet-address"`
	LACPInterfaceOffset     int         `json:"lacp-interface-offset"`
	LACPSystemMAC           string      `json:"lacp-system-mac"`
	LeafGroup               string      `json:"leaf-group"`
	ModelNumberDescription  string      `json:"model-number-description"`
	Name                    string      `json:"name"`
	SerialNumberDescription string      `json:"serial-number-description"`
	Shutdown                bool        `json:"shutdown"`
}

type Switches []Switch
