package hostcheck

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/marcelmiguel/ping"
)

type RequestorError struct {
	msg string
}

func (err RequestorError) Error() string {
	return err.msg
}

type Requestor interface {
	DoICMP(host string) (RequestResult, error)
	DoTCP(host string) (RequestResult, error)
	DoUDP(host string) (RequestResult, error)
	DoHTTP(host string) (RequestResult, error)
	DoHTTPS(host string) (RequestResult, error)
}

type BaseRequestor struct{}

var emptyResult = RequestResult{Timing: time.Nanosecond, ResponseStatus: ResponseStatus{IsSuccess: false, StatusCode: -1}}

func isPortOpen(host string, port Port, protocol Protocol, timeout time.Duration) (bool, time.Duration) {
	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))

	start_at := time.Now()
	conn, err := net.DialTimeout(string(protocol), address, timeout)
	if err != nil {
		return false, 0
	}
	conn.Close()
	return true, time.Since(start_at)
}

func doTransportProtocolCheck(host string, protocol Protocol) (RequestResult, error) {
	for _, port := range TransportProtocolPorts[protocol] {
		is_open, timing := isPortOpen(host, port, protocol, 10*time.Second)
		if is_open {
			return RequestResult{timing, ResponseStatus{true, 0}},
				nil
		}
	}

	return RequestResult{time.Duration(-1), ResponseStatus{false, 0}},
		RequestorError{fmt.Sprintf("chosen ports are closed for %s", protocol)}
}

func doTransferProtocolCheck(host string, protocol Protocol) (RequestResult, error) {
	prefix := TransferProtocolPrefix[protocol]
	url := fmt.Sprintf("%s%s", prefix, host)
	start_at := time.Now()
	resp, err := http.Get(url)
	return RequestResult{time.Since(start_at), ResponseStatus{err == err, resp.StatusCode}}, nil
}

func (req BaseRequestor) DoICMP(host string) (RequestResult, error) {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		return emptyResult, err
	}

	pinger.Count = 1
	err = pinger.Run()

	if err != nil {
		return emptyResult, err
	}

	stats := pinger.Statistics()
	return RequestResult{stats.AvgRtt, ResponseStatus{true, 0}}, nil
}

func (req BaseRequestor) DoTCP(host string) (RequestResult, error) {
	return doTransportProtocolCheck(host, TCP)
}

func (req BaseRequestor) DoUDP(host string) (RequestResult, error) {
	return doTransportProtocolCheck(host, UDP)
}

func (req BaseRequestor) DoHTTP(host string) (RequestResult, error) {
	return doTransferProtocolCheck(host, HTTP)
}

func (req BaseRequestor) DoHTTPS(host string) (RequestResult, error) {
	return doTransferProtocolCheck(host, HTTPS)
}
