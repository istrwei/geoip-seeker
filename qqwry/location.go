package qqwry

import (
	"net"
	"strings"
)

type Location struct {
	IP      net.IP `json:"ip"`
	BeginIP net.IP `json:"begin_ip"`
	EndIP   net.IP `json:"end_ip"`
	Country string `json:"country"`
	Area    string `json:"area"`
}

func (location *Location) String() string {
	items := []string{
		location.BeginIP.String(),
		location.EndIP.String(),
		location.Country,
		location.Area,
	}
	return strings.Join(items, "\t")
}
