package hostcheck

type Protocol int

const (
	ICMP Protocol = iota
	TCP
	UDP
	HTTP
	HTTPS
)

var Protocols = []Protocol{ICMP, TCP, UDP, HTTP, HTTPS}

type Port int

var ProtocolStr = map[Protocol]string{
	ICMP:  "ICMP",
	TCP:   "TCP",
	UDP:   "UDP",
	HTTP:  "HTTP",
	HTTPS: "HTTPS",
}
var TransportProtocolStr = map[Protocol]string{
	TCP: "tcp",
	UDP: "udp",
}

var TransportProtocolPorts = map[Protocol][]Port{
	TCP: {80, 443},
	UDP: {443},
}

var TransferProtocolPrefix = map[Protocol]string{
	HTTP:  "http://",
	HTTPS: "https://",
}
