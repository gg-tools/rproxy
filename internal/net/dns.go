package net

import (
	"context"
	"net"
	"strings"
)

func ConfigDefaultResolver(dnsServer string) {
	if dnsServer == "" {
		return // use default
	}

	s := dnsServer
	if !strings.Contains(s, ":") {
		s += ":53"
	}
	net.DefaultResolver = &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return net.Dial("udp", s)
		},
	}
}
