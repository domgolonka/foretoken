package providers

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
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
		c.iplist = make([]string, 0, 0)
		c.sublist = make([]string, 0, 0)
	}

	regexp_ip := "((?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?))"
	regexp_subnet := regexp_ip + "\\/(3[0-1]|[1-2][0-9]|[1-9])"

	spamhaus := Feed{"spamhaus", "https://www.spamhaus.org/drop/drop.txt",
		10, []FeedAnalyzer{{3, "^" + regexp_subnet + ".*"},
			{3, "^" + regexp_ip + ".*"}}}
	firehol := Feed{"firehol", "https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/firehol_level1.netset",
		10, []FeedAnalyzer{{3, "^" + regexp_subnet + ".*"},
			{3, "^" + regexp_ip + ".*"}}}
	alienvault_com := Feed{"alienvault.com", "https://reputation.alienvault.com/reputation.generic",
		10, []FeedAnalyzer{{1, "^" + regexp_ip + " # Scanning Host.*"},
			{3, "^" + regexp_ip + " # Malicious Host.*"}}}
	badips_com := Feed{"badips.com", "https://www.badips.com/get/list/any/2?age=7d", 10,
		[]FeedAnalyzer{{3, "^" + regexp_ip + ".*"}}}
	team_cymru_org := Feed{"team-cymru.org", "https://www.team-cymru.org/Services/Bogons/fullbogons-ipv4.txt", 10,
		[]FeedAnalyzer{{3, "^" + regexp_subnet + ".*"}}}
	stopforumspam_com := Feed{"stopforumspam.com", "https://www.stopforumspam.com/downloads/toxic_ip_cidr.txt", 10,
		[]FeedAnalyzer{{2, "^" + regexp_subnet + ".*"}}}
	greensnow_co := Feed{"greensnow.co", "https://blocklist.greensnow.co/greensnow.txt",
		10, []FeedAnalyzer{{2, "^" + regexp_ip + ".*"}}}
	binarydefense_com := Feed{"binarydefense.com", "https://www.binarydefense.com/banlist.txt",
		10, []FeedAnalyzer{{1, "^" + regexp_ip + ".*"}}}
	haleys_org_ssh := Feed{"the-haleys.org", "http://charles.the-haleys.org/ssh_dico_attack_with_timestamps.php?days=7",
		10, []FeedAnalyzer{{1, "^ALL : ((?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)).*"}}}
	haleys_org_wp := Feed{"the-haleys.org", "http://charles.the-haleys.org/wp_attack_with_timestamps.php?days=7",
		10, []FeedAnalyzer{{1, "^" + regexp_ip + ".*"}}}
	haleys_org_smtp := Feed{"the-haleys.org", "http://charles.the-haleys.org/smtp_dico_attack_with_timestamps.php?days=7",
		10, []FeedAnalyzer{{1, "^" + regexp_ip + ".*"}}}
	bambenekconsulting_com := Feed{"bambenekconsulting.com", "http://osint.bambenekconsulting.com/feeds/c2-ipmasterlist-high.txt",
		10, []FeedAnalyzer{{3, "^" + regexp_ip + ".*"}}}
	blocklist_de := Feed{"blocklist.de", "http://lists.blocklist.de/lists/all.txt", 10,
		[]FeedAnalyzer{{3, "^" + regexp_ip + ".*"}}}
	botscout := Feed{"botscout", "https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/botscout_1d.ipset",
		10, []FeedAnalyzer{{3, "^" + regexp_ip + ".*"}}}
	bruteforceblocker := Feed{"bruteforceblocker", "http://danger.rulez.sk/projects/bruteforceblocker/blist.php",
		10, []FeedAnalyzer{{3, "^" + regexp_ip + ".*"}}}
	cinsscore_com := Feed{"cinsscore.com", "http://cinsscore.com/list/ci-badguys.txt", 10,
		[]FeedAnalyzer{{3, "^" + regexp_ip + ".*"}}}
	cruzit := Feed{"cruzit", "https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/cruzit_web_attacks.ipset",
		10, []FeedAnalyzer{{3, "^" + regexp_ip + ".*"}}}
	dshield_org := Feed{"dshield.org", "http://feeds.dshield.org/top10-2.txt", 10,
		[]FeedAnalyzer{{3, "^" + regexp_ip + ".*"}}}
	emergingthreats_net := Feed{"emergingthreats.net", "http://rules.emergingthreats.net/open/suricata/rules/compromised-ips.txt",
		10, []FeedAnalyzer{{3, "^" + regexp_ip + ".*"}}}
	malwaredomainlist := Feed{"malwaredomainlist", "https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/malwaredomainlist.ipset",
		10, []FeedAnalyzer{{3, "^" + regexp_ip + ".*"}}}
	myip := Feed{"myip", "https://myip.ms/files/blacklist/htaccess/latest_blacklist.txt", 10,
		[]FeedAnalyzer{{3, "^deny from ((?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)).*"}}}

	sslbl := Feed{"sslbl", "https://sslbl.abuse.ch/blacklist/sslipblacklist_aggressive.txt",
		10, []FeedAnalyzer{{3, "^" + regexp_ip + ".*"}}}
	rutgers_edu := Feed{"rutgers.edu", "https://report.cs.rutgers.edu/DROP/attackers",
		10, []FeedAnalyzer{{3, "^" + regexp_ip + ".*"}}}
	sblam_com := Feed{"sblam.com", "http://sblam.com/blacklist.txt",
		10, []FeedAnalyzer{{1, "^" + regexp_ip + ".*"}}}
	talosintelligence_com := Feed{"talosintelligence.com", "http://www.talosintelligence.com/feeds/ip-filter.blf",
		10, []FeedAnalyzer{{3, "^" + regexp_ip + ".*"}}}
	turris_cz := Feed{"turris.cz", "https://www.turris.cz/greylist-data/greylist-latest.csv",
		10, []FeedAnalyzer{{3, "^" + regexp_ip + ".*"}}}
	feodotracker := Feed{"feodotracker", "https://feodotracker.abuse.ch/downloads/ipblocklist_aggressive.txt",
		10, []FeedAnalyzer{{3, "^" + regexp_ip + ".*"}}}
	var active_feeds = []Feed{team_cymru_org, stopforumspam_com, greensnow_co, binarydefense_com, haleys_org_ssh,
		haleys_org_wp, haleys_org_smtp, spamhaus, firehol, alienvault_com, badips_com, bambenekconsulting_com, blocklist_de, botscout,
		bruteforceblocker, cinsscore_com, cruzit, dshield_org, emergingthreats_net, feodotracker, malwaredomainlist,
		myip, sslbl,
		rutgers_edu, sblam_com, talosintelligence_com,
		turris_cz, feodotracker}
	ips := make(map[string]IPAnalysis)
	subnets := make(map[string]SUBNETAnalysis)
	for _, active_feed := range active_feeds {
		c.logger.Printf("[INFO] Importing data feed %s\n", active_feed.Name)
		feed_results_ips, feed_results_subnets, err := active_feed.Fetch()

		if err == nil {
			for k, e := range feed_results_ips {
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
			for k, e := range feed_results_subnets {
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
			c.logger.Printf("[INFO] Imported %d ips and %d subnets from data feed %s\n", len(feed_results_ips),
				len(feed_results_subnets), active_feed.Name)
		} else {
			c.logger.Printf("[ERROR] Importing data feed %s\n failed : %s", active_feed.Name, err)
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
