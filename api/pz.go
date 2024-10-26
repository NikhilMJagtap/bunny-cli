package api

import (
	"fmt"
	"net"

	"github.com/NikhilMJagtap/bunny-cli/client"
)

type ListPullZoneOpts struct {
	Page               uint32 `json:"page"`
	PerPage            uint32 `json:"perPage"`
	Search             string `json:"search"`
	IncludeCertificate bool   `json:"includeCertificate"`
}

func ListPullZones(b *client.BunnyClient, options *ListPullZoneOpts) (interface{}, error) {
	return b.Get("/pullzone/", GetQueryParamsFromOptions(options))
}

func GetPullZone(b *client.BunnyClient, pullZoneId uint64) (interface{}, error) {
	return b.Get(fmt.Sprintf("/pullzone/%d/", pullZoneId), nil)
}

type BlockIPData struct {
	IP string `json:"BlockedIp"`
}

func BlockIP(b *client.BunnyClient, pullZoneId uint64, ip net.IP) (interface{}, error) {
	data := BlockIPData{
		IP: ip.String(),
	}
	return b.Post(fmt.Sprintf("/pullzone/%d/addBlockedIp/", pullZoneId), data)
}

func UnblockIP(b *client.BunnyClient, pullZoneId uint64, ip net.IP) (interface{}, error) {
	data := BlockIPData{
		IP: ip.String(),
	}
	return b.Post(fmt.Sprintf("/pullzone/%d/removeBlockedIp/", pullZoneId), data)
}
