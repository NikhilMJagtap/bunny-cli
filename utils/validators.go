package utils

import (
	"errors"
	"net"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

func ValidateIP(ip string) (string, error) {
	if strings.Contains(ip, "*") {

		ipv4Pattern := `^((\d{1,2}|1\d{2}|2[0-4]\d|25[0-5])|\*)(\.((\d{1,2}|1\d{2}|2[0-4]\d|25[0-5])|\*)){3}$`
		if matched, _ := regexp.MatchString(ipv4Pattern, ip); matched {
			return ip, nil
		}

		ipv6Pattern := `^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$`
		if matched, _ := regexp.MatchString(ipv6Pattern, ip); matched {
			return ip, nil
		}

		return "", errors.New("Invalid IP pattern. Received " + ip)
	}

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return "", errors.New("Invalid IP address. Received " + ip)
	}
	return ip, nil
}

func AddRemoveValidator(cmd *cobra.Command, args []string) error {
	action := args[0]
	if action != "add" && action != "remove" {
		return errors.New("invalid action: must be either 'add' or 'remove'")
	}
	return nil
}
