// Copyright 2023 Christopher Briscoe.  All rights reserved.
package net

import (
	"net"
	"net/http"
)

// GetIP will first try to get proxy IP headers in case the server is running behind a
// reverse proxy.  Otherwise, just use the http.Request.RemoteAddr with the port stripped.
func GetIP(req *http.Request) string {
	ip := req.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}

	ip = req.Header.Get("X-Forwarded-For")
	if ip != "" {
		return ip
	}

	ip, _, _ = net.SplitHostPort(req.RemoteAddr)
	return ip
}
