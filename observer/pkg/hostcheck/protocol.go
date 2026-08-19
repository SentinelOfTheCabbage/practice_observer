package hostcheck

type Protocol string

const (
	ICMP  Protocol = "ICMP"
	TCP   Protocol = "TCP"
	UDP   Protocol = "UDP"
	HTTP  Protocol = "HTTP"
	HTTPS Protocol = "HTTPS"
)

var Protocols = []Protocol{ICMP, TCP, UDP, HTTP, HTTPS}

type Port int

var TransportProtocolPorts = map[Protocol][]Port{
	TCP: {80, 443},
	UDP: {443},
}

var TransferProtocolPrefix = map[Protocol]string{
	HTTP:  "http://",
	HTTPS: "https://",
}
