package api

import (
	"fmt"

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

func BlockIP(b *client.BunnyClient, pullZoneId uint64, ip string) (interface{}, error) {
	data := BlockIPData{
		IP: ip,
	}
	return b.Post(fmt.Sprintf("/pullzone/%d/addBlockedIp/", pullZoneId), data)
}

func UnblockIP(b *client.BunnyClient, pullZoneId uint64, ip string) (interface{}, error) {
	data := BlockIPData{
		IP: ip,
	}
	return b.Post(fmt.Sprintf("/pullzone/%d/removeBlockedIp/", pullZoneId), data)
}

type AllowedReferrerData struct {
	Referrer string `json:"Hostname"`
}

func AddAllowedReferrer(b *client.BunnyClient, pullZoneId uint64, referrer string) (interface{}, error) {
	data := AllowedReferrerData{
		Referrer: referrer,
	}
	return b.Post(fmt.Sprintf("/pullzone/%d/addAllowedReferrer/", pullZoneId), data)
}

func RemoveAllowedReferrer(b *client.BunnyClient, pullZoneId uint64, referrer string) (interface{}, error) {
	data := AllowedReferrerData{
		Referrer: referrer,
	}
	return b.Post(fmt.Sprintf("/pullzone/%d/removeAllowedReferrer/", pullZoneId), data)
}

type BlockedReferrerData struct {
	Referrer string `json:"Hostname"`
}

func AddBlockedReferrer(b *client.BunnyClient, pullZoneId uint64, referrer string) (interface{}, error) {
	data := BlockedReferrerData{
		Referrer: referrer,
	}
	return b.Post(fmt.Sprintf("/pullzone/%d/addBlockedReferrer/", pullZoneId), data)
}

func RemoveBlockedReferrer(b *client.BunnyClient, pullZoneId uint64, referrer string) (interface{}, error) {
	data := BlockedReferrerData{
		Referrer: referrer,
	}
	return b.Post(fmt.Sprintf("/pullzone/%d/removeBlockedReferrer/", pullZoneId), data)
}

type HostnameData struct {
	Hostname string `json:"Hostname"`
}

func AddHostname(b *client.BunnyClient, pullZoneId uint64, hostname string) (interface{}, error) {
	data := HostnameData{
		Hostname: hostname,
	}
	return b.Post(fmt.Sprintf("/pullzone/%d/addHostname/", pullZoneId), data)
}

func RemoveHostname(b *client.BunnyClient, pullZoneId uint64, hostname string) (interface{}, error) {
	data := HostnameData{
		Hostname: hostname,
	}
	return b.Delete(fmt.Sprintf("/pullzone/%d/removeHostname/", pullZoneId), data)
}

type ForceSSLData struct {
	HostnameData
	ForceSSL bool `json:"ForceSSL"`
}

func ForceSSL(b *client.BunnyClient, pullZoneId uint64, hostname string, enable bool) (interface{}, error) {
	data := ForceSSLData{
		HostnameData: HostnameData{
			Hostname: hostname,
		},
		ForceSSL: enable,
	}
	return b.Post(fmt.Sprintf("/pullzone/%d/setForceSSL/", pullZoneId), data)
}
