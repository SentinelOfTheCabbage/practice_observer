package hostcheck

func DoICMP(host string, req Requestor) SingleResult {
	request_result, _ := req.DoICMP(host)
	return SingleResult{
		Host:          host,
		Protocol:      ICMP,
		RequestResult: request_result,
	}
}

func DoTCP(host string, req Requestor) SingleResult {
	request_result, _ := req.DoTCP(host)
	return SingleResult{
		Host:          host,
		Protocol:      TCP,
		RequestResult: request_result,
	}
}

func DoUDP(host string, req Requestor) SingleResult {
	request_result, _ := req.DoUDP(host)
	return SingleResult{
		Host:          host,
		Protocol:      UDP,
		RequestResult: request_result,
	}
}

func DoHTTP(host string, req Requestor) SingleResult {
	request_result, _ := req.DoHTTP(host)
	return SingleResult{
		Host:          host,
		Protocol:      HTTP,
		RequestResult: request_result,
	}
}

func DoHTTPS(host string, req Requestor) SingleResult {
	request_result, _ := req.DoHTTPS(host)
	return SingleResult{
		Host:          host,
		Protocol:      HTTPS,
		RequestResult: request_result,
	}
}
