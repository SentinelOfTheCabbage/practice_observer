package main

import (
	"fmt"
	"observer/pkg/dependencies"
	"observer/pkg/env"
	"observer/pkg/hostcheck"
	"strings"
	"syscall"
)

func getHosts() []string {
	env.GetEnv()
	urls_to_check, err := syscall.Getenv("URLS_TO_CHECK")
	if err {
		panic("Can't load urls to check list!")
	}
	return strings.Split(urls_to_check, ",")
}

func doRequest(host string, protocol hostcheck.Protocol, req hostcheck.Requestor) hostcheck.SingleResult {
	switch protocol {
	case hostcheck.ICMP:
		return hostcheck.DoICMP(host, req)
	case hostcheck.TCP:
		return hostcheck.DoTCP(host, req)
	case hostcheck.UDP:
		return hostcheck.DoUDP(host, req)
	case hostcheck.HTTP:
		return hostcheck.DoHTTP(host, req)
	case hostcheck.HTTPS:
		return hostcheck.DoHTTPS(host, req)
	default:
		panic(fmt.Sprintf("Unknown protocol: %d", protocol))
	}
}

func SingleCheck(host string, protocol hostcheck.Protocol, deps dependencies.DepsInterface) {
	result := doRequest(host, protocol, deps.GetRequestor())
	deps.SendToKafka(result)
}

func RunCheck(host string, protocol hostcheck.Protocol, deps dependencies.DepsInterface) {
	for {
		go SingleCheck(host, protocol, deps)
	}
}

func RunChecks(deps dependencies.DepsInterface) {
	hosts := getHosts()
	for _, host := range hosts {
		for _, protocol := range hostcheck.Protocols {
			go RunCheck(host, protocol, deps)
		}
	}
}
