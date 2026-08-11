package hostcheck

import "time"

type ResponseStatus struct {
	IsSuccess  bool
	StatusCode int
}

type RequestResult struct {
	Timing         time.Duration
	ResponseStatus ResponseStatus
}

type SingleResult struct {
	Host          string
	Protocol      Protocol
	RequestResult RequestResult
}

type HostResult struct {
	Host           string
	RequestResults map[Protocol]RequestResult
}
