package runcheck

import (
	"observer/pkg/dependencies"
	"observer/pkg/hostcheck"
	"time"
)

func SingleCheck(host string, protocol hostcheck.Protocol, deps dependencies.DepsInterface) {
	result := hostcheck.DoRequest(host, protocol, deps.GetRequestor())
	deps.SendToKafka(result)
}

func RunCheck(host string, protocol hostcheck.Protocol, settings hostcheck.Settings, deps dependencies.DepsInterface) {
	for {
		SingleCheck(host, protocol, deps)
		time.Sleep(time.Millisecond * time.Duration(settings.SleepMs))
	}
}

func RunChecks(deps dependencies.DepsInterface) {
	settings := hostcheck.GetSettings()
	for _, host := range settings.UrlsToCheck {
		for _, protocol := range settings.ProtocolsToCheck {
			go RunCheck(host, protocol, settings, deps)
		}
	}
}
