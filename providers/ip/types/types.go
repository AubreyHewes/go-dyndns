package types

import "net"

type Provider interface {
	GetIPs() ([]IPResult, error)
}

type IPResult struct {
	Type string
	IP   net.IP
}
