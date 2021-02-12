package providers

type GoldenFrog struct {
}

func NewGoldenFrog() *OpenVpn {
	return &OpenVpn{}
}
func (*GoldenFrog) Name() string {
	return "openvpn"
}

func (*GoldenFrog) Format(src string) ([]string, error) {
	hosts := []string{}

	return hosts, nil
}
