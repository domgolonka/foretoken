package providers

import (
	"bytes"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type TxtDomains struct {
	iplist     []string
	sublist    []string
	logger     logrus.FieldLogger
	lastUpdate time.Time
}

func NewTxtDomains(logger logrus.FieldLogger) *TxtDomains {
	logger.Debug("starting TxtDomains")
	return &TxtDomains{
		logger: logger,
	}
}
func (*TxtDomains) Name() string {
	return "text_domain"
}

func (c *TxtDomains) Load(body []byte) ([]string, []string, error) {
	iplist := []string{}
	sublist := []string{}

	// don't need to update this more than once a day!
	if time.Now().Unix() >= c.lastUpdate.Unix()+(82800) {
		c.iplist = make([]string, 0)
		c.sublist = make([]string, 0)
	}

	regexpIP := "((?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?))"
	regexpSubnet := regexpIP + "\\/(3[0-1]|[1-2][0-9]|[1-9])"

	spamhaus := Feed{"spamhaus", "https://www.spamhaus.org/drop/drop.txt",
		10, []FeedAnalyzer{{3, "^" + regexpSubnet + ".*"},
			{3, "^" + regexpIP + ".*"}}}
	firehol := Feed{"firehol", "https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/firehol_level1.netset",
		10, []FeedAnalyzer{{3, "^" + regexpSubnet + ".*"},
			{3, "^" + regexpIP + ".*"}}}
	alienvaultCom := Feed{"alienvault.com", "https://reputation.alienvault.com/reputation.generic",
		10, []FeedAnalyzer{{1, "^" + regexpIP + " # Scanning Host.*"},
			{3, "^" + regexpIP + " # Malicious Host.*"}}}
	badipsCom := Feed{"badips.com", "https://www.badips.com/get/list/any/2?age=7d", 10,
		[]FeedAnalyzer{{3, "^" + regexpIP + ".*"}}}
	teamCymruOrg := Feed{"team-cymru.org", "https://www.team-cymru.org/Services/Bogons/fullbogons-ipv4.txt", 10,
		[]FeedAnalyzer{{3, "^" + regexpSubnet + ".*"}}}
	stopforumspamCom := Feed{"stopforumspam.com", "https://www.stopforumspam.com/downloads/toxic_ip_cidr.txt", 10,
		[]FeedAnalyzer{{2, "^" + regexpSubnet + ".*"}}}
	greensnowCo := Feed{"greensnow.co", "https://blocklist.greensnow.co/greensnow.txt",
		10, []FeedAnalyzer{{2, "^" + regexpIP + ".*"}}}
	binarydefenseCom := Feed{"binarydefense.com", "https://www.binarydefense.com/banlist.txt",
		10, []FeedAnalyzer{{1, "^" + regexpIP + ".*"}}}
	haleysOrgSSH := Feed{"the-haleys.org", "http://charles.the-haleys.org/ssh_dico_attack_with_timestamps.php?days=7",
		10, []FeedAnalyzer{{1, "^ALL : ((?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)).*"}}}
	haleysOrgWp := Feed{"the-haleys.org", "http://charles.the-haleys.org/wp_attack_with_timestamps.php?days=7",
		10, []FeedAnalyzer{{1, "^" + regexpIP + ".*"}}}
	haleysOrgSMTP := Feed{"the-haleys.org", "http://charles.the-haleys.org/smtp_dico_attack_with_timestamps.php?days=7",
		10, []FeedAnalyzer{{1, "^" + regexpIP + ".*"}}}
	blocklistDe := Feed{"blocklist.de", "http://lists.blocklist.de/lists/all.txt", 10,
		[]FeedAnalyzer{{3, "^" + regexpIP + ".*"}}}
	botscout := Feed{"botscout", "https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/botscout_1d.ipset",
		10, []FeedAnalyzer{{3, "^" + regexpIP + ".*"}}}
	bruteforceblocker := Feed{"bruteforceblocker", "http://danger.rulez.sk/projects/bruteforceblocker/blist.php",
		10, []FeedAnalyzer{{3, "^" + regexpIP + ".*"}}}
	cinsscoreCom := Feed{"cinsscore.com", "http://cinsscore.com/list/ci-badguys.txt", 10,
		[]FeedAnalyzer{{3, "^" + regexpIP + ".*"}}}
	cruzit := Feed{"cruzit", "https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/cruzit_web_attacks.ipset",
		10, []FeedAnalyzer{{3, "^" + regexpIP + ".*"}}}
	dshieldOrg := Feed{"dshield.org", "http://feeds.dshield.org/top10-2.txt", 10,
		[]FeedAnalyzer{{3, "^" + regexpIP + ".*"}}}
	emergingthreatsNet := Feed{"emergingthreats.net", "http://rules.emergingthreats.net/open/suricata/rules/compromised-ips.txt",
		10, []FeedAnalyzer{{3, "^" + regexpIP + ".*"}}}
	malwaredomainlist := Feed{"malwaredomainlist", "https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/malwaredomainlist.ipset",
		10, []FeedAnalyzer{{3, "^" + regexpIP + ".*"}}}
	myip := Feed{"myip", "https://myip.ms/files/blacklist/htaccess/latest_blacklist.txt", 10,
		[]FeedAnalyzer{{3, "^deny from ((?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)).*"}}}

	sslbl := Feed{"sslbl", "https://sslbl.abuse.ch/blacklist/sslipblacklist_aggressive.txt",
		10, []FeedAnalyzer{{3, "^" + regexpIP + ".*"}}}
	rutgersEdu := Feed{"rutgers.edu", "http://report.cs.rutgers.edu/DROP/attackers",
		10, []FeedAnalyzer{{3, "^" + regexpIP + ".*"}}}
	sblamCom := Feed{"sblam.com", "http://sblam.com/blacklist.txt",
		10, []FeedAnalyzer{{1, "^" + regexpIP + ".*"}}}
	talosintelligenceCom := Feed{"talosintelligence.com", "http://www.talosintelligence.com/feeds/ip-filter.blf",
		10, []FeedAnalyzer{{3, "^" + regexpIP + ".*"}}}
	turrisCz := Feed{"turris.cz", "https://www.turris.cz/greylist-data/greylist-latest.csv",
		10, []FeedAnalyzer{{3, "^" + regexpIP + ".*"}}}
	feodotracker := Feed{"feodotracker", "https://feodotracker.abuse.ch/downloads/ipblocklist_aggressive.txt",
		10, []FeedAnalyzer{{3, "^" + regexpIP + ".*"}}}
	fireholdabusers := Feed{"boyscout1d", "https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/firehol_abusers_1d.netset",
		10, []FeedAnalyzer{{1, "^" + regexpIP + ".*"}}}
	var activeFeeds = []Feed{teamCymruOrg, stopforumspamCom, greensnowCo, binarydefenseCom, haleysOrgSSH,
		haleysOrgWp, haleysOrgSMTP, spamhaus, firehol, alienvaultCom, badipsCom, blocklistDe, botscout,
		bruteforceblocker, cinsscoreCom, cruzit, dshieldOrg, emergingthreatsNet, feodotracker, malwaredomainlist,
		myip, sslbl, rutgersEdu, sblamCom, talosintelligenceCom,
		turrisCz, feodotracker, fireholdabusers}
	ips := make(map[string]IPAnalysis)
	subnets := make(map[string]SUBNETAnalysis)
	for _, activeFeed := range activeFeeds {
		c.logger.Printf("[INFO] Importing data feed %s\n", activeFeed.Name)
		feedResultsIPs, feedResultsSubnets, err := activeFeed.Fetch()

		if err == nil {
			for k, e := range feedResultsIPs {
				if _, ok := ips[k]; ok {
					ip := ips[k]
					iplist = append(iplist, ip.IP)
					ip.Score = ip.Score + e.Score
					ip.Lists = append(ip.Lists, e.Lists[0])
					ips[k] = ip
				} else {
					ips[k] = e
				}
			}
			for k, e := range feedResultsSubnets {
				if _, ok := subnets[k]; ok {
					subnet := subnets[k]
					sublist = append(sublist, subnet.SUBNET)
					subnet.Score = subnet.Score + e.Score
					subnet.Lists = append(subnet.Lists, e.Lists[0])
					subnets[k] = subnet
				} else {
					subnets[k] = e
				}
			}
			c.logger.Printf("[INFO] Imported %d ips and %d subnets from data feed %s\n", len(feedResultsIPs),
				len(feedResultsSubnets), activeFeed.Name)
		} else {
			c.logger.Printf("[ERROR] Importing data feed %s\n failed : %s", activeFeed.Name, err)
		}
	}

	//if len(c.hosts) != 0 {
	//	return c.hosts, nil
	//}
	//allbody := make([]byte, 0, len(domains))
	//if body == nil {
	//	var err error
	//	for i := 0; i < len(domains); i++ {
	//		c.logger.Debug(domains[i])
	//		if body, err = c.MakeRequest(domains[i]); err == nil {
	//			allbody = append(allbody, body...)
	//		}
	//
	//	}
	//}
	//
	//c.hosts = strings.Split(string(allbody), "\n")

	c.lastUpdate = time.Now()
	c.iplist = iplist
	c.sublist = sublist
	return iplist, sublist, nil

}
func (c *TxtDomains) MakeRequest(urllist string) ([]byte, error) {
	var body bytes.Buffer

	req, err := http.NewRequest(http.MethodGet, urllist, nil)
	if err != nil {
		return nil, err
	}

	var client = NewClient()

	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	return body.Bytes(), err
}

func (c *TxtDomains) List() ([]string, []string, error) {
	return c.Load(nil)
}
