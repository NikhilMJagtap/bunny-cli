package api

import (
    "github.com/NikhilMJagtap/bunny-cli/client"
)

func ListRegions(b *client.BunnyClient) (interface{}, error) {
    return b.Get("/region/", nil)
}
