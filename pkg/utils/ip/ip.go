package ip

import (
	"net"
	"regexp"
	"strconv"
	"strings"
)

func ParseIPs(body []byte, expression []string) []string {
	splitup := strings.Split(string(body), "\n")
	ipv4 := make([]string, 0, len(splitup))
	for i := 0; i < len(splitup); i++ {
		for _, exp := range expression {
			isIP := parseIps(splitup[i], exp)

			if isIP {
				ipv4 = append(ipv4, splitup[i])
			}
		}

	}
	return ipv4
}

// Check for IP:PORT
func parseIps(str string, expression string) bool {
	if str == "" {
		return false
	}
	ipv4WithPortCidrRegex := regexp.MustCompile(`` + expression + ``)
	return ipv4WithPortCidrRegex.MatchString(str)

}

func ParseSubnet(ip, ipSub string, subnet byte) bool {
	ipstr := ipSub + "/" + strconv.Itoa(int(subnet))
	singleIP := ip + "/32"
	_, ipnetA, _ := net.ParseCIDR(ipstr)
	ipB, _, _ := net.ParseCIDR(singleIP)
	return ipnetA.Contains(ipB)
}

/*
func ParseIps(body []byte) ([]string, error) {
	splitup := strings.Split(string(body), "\n")
	ipv4 := make([]string, 0, len(splitup))
	for i := 0; i < len(splitup); i++ {
		isIPv4 := parseIpv4AndPort(splitup[i])

		if isIPv4 {
			ipv4 = append(ipv4, splitup[i])
		}
	}

	return ipv4, nil

}

// Check for IP:PORT
func parseIpv4AndPort(str string) (m bool) {
	if str == "" {
		return false
	}
	// first check for port and IP
	ipv4AndPortWithCidrRegex := regexp.MustCompile(`(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)(?:\/([1-3]?\d))?:\d{1,5}`)
	if ipv4AndPortWithCidrRegex.MatchString(str) {
		return true
	}

	ipv4WithPortCidrRegex := regexp.MustCompile(`(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)`)
	return ipv4WithPortCidrRegex.MatchString(str)

}
*/
