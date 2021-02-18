package ip

import (
	"regexp"
	"strings"
)

func ParseIps(body []byte) ([]string, error) {
	splitup := strings.Split(string(body), "\n")
	ipv4 := make([]string, 0, len(splitup))
	for i := 0; i < len(splitup); i++ {
		isIPv4 := ParseIpv4AndPort(splitup[i])

		if isIPv4 {
			ipv4 = append(ipv4, splitup[i])
		}
	}

	return ipv4, nil

}

// Check for IP:PORT
func ParseIpv4AndPort(str string) (m bool) {
	if str == "" {
		return false
	}
	ipv4WithCidrRegex := regexp.MustCompile(`(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)(?:\/([1-3]?\d))?:\d{1,5}`)
	return ipv4WithCidrRegex.MatchString(str)
}
