package hostcheck

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3" // Correct import
)

type Settings struct {
	UrlsToCheck      []string   `yaml:"urls"`
	ProtocolsToCheck []Protocol `yaml:"protocols"`
	SleepMs          int        `yaml:"sleep_ms"`
}

func GetSettings() Settings {
	data, err := os.ReadFile("../../configs/config.yaml")
	if err != nil {
		panic(err)
	}
	var result Settings

	err = yaml.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}
	return result
}

func DoRequest(host string, protocol Protocol, req Requestor) SingleResult {
	switch protocol {
	case ICMP:
		return DoICMP(host, req)
	case TCP:
		return DoTCP(host, req)
	case UDP:
		return DoUDP(host, req)
	case HTTP:
		return DoHTTP(host, req)
	case HTTPS:
		return DoHTTPS(host, req)
	default:
		panic(fmt.Sprintf("Unsupported protocol: %s", protocol))
	}
}
