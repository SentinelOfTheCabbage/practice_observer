package main

import (
	"observer/pkg/dependencies"
	"observer/pkg/runcheck"
)

func main() {
	var deps = dependencies.Dependencies{}
	runcheck.RunChecks(deps)

	select {}
}
