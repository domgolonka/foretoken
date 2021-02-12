package ip

import (
	"regexp"
	"strings"
)

func ParseIps(body []byte) ([]string, error) {

	splitup := strings.Split(string(body), "\n")
	ipv4 := make([]string, 0, len(splitup))
	for i := 0; i < len(splitup); i++ {
		isipv4 := ParseIpv4(splitup[i])

		if isipv4 {
			ipv4 = append(ipv4, splitup[i])
		}
	}

	return ipv4, nil

}

func ParseIpv4(str string) (m bool) {
	if str == "" {
		return false
	}

	// Regex didn't work when I tried to compress it.. so I guess we get to use the expanded version. Written by hand
	ipv4WithCidrRegex := regexp.MustCompile(`(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)(?:\/([1-3]?\d))?`)

	return ipv4WithCidrRegex.MatchString(str)
}
