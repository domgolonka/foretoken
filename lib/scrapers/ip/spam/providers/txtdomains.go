package providers

import (
	"bytes"
	"net/http"
	"time"

	"github.com/domgolonka/threatdefender/app/models"

	"github.com/domgolonka/threatdefender/app/entity"

	"github.com/sirupsen/logrus"
)

type TxtDomains struct {
	iplist     []models.Spam
	sublist    []models.Spam
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

func (c *TxtDomains) Load(body []byte) ([]models.Spam, []models.Spam, error) {

	// don't need to update this more than once a day!
	if time.Now().Unix() >= c.lastUpdate.Unix()+(82800) {
		c.iplist = make([]models.Spam, 0)
		c.sublist = make([]models.Spam, 0)
	}

	//regexpIP := "((?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?))"
	//regexpSubnet := regexpIP + "\\/(3[0-1]|[1-2][0-9]|[1-9])"

	//spamhaus := entity.Feed{Name: "spamhaus", URL: "https://www.spamhaus.org/drop/drop.txt",
	//	Timeout: 10, FeedAnalyzers: []entity.FeedAnalyzer{{Score: 3, Expression: "^" + regexpSubnet + ".*"},
	//		{Score: 3, Expression: "^" + regexpIP + ".*"}}}
	//firehol := entity.Feed{Name: "firehol", URL: "https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/firehol_level1.netset",
	//	Timeout: 10, FeedAnalyzers: []entity.FeedAnalyzer{{Score: 3, Expression: "^" + regexpSubnet + ".*"},
	//		{Score: 3, Expression: "^" + regexpIP + ".*"}}}
	//alienvaultCom := entity.Feed{Name: "alienvault.com", URL: "https://reputation.alienvault.com/reputation.generic",
	//	Timeout: 10, FeedAnalyzers: []entity.FeedAnalyzer{{Score: 1, Expression: "^" + regexpIP + " # Scanning Host.*"},
	//		{Score: 3, Expression: "^" + regexpIP + " # Malicious Host.*"}}}
	//badipsCom := entity.Feed{Name: "badips.com", URL: "https://www.badips.com/get/list/any/2?age=7d", Timeout: 10,
	//	FeedAnalyzers: []entity.FeedAnalyzer{{Score: 3, Expression: "^" + regexpIP + ".*"}}}
	//teamCymruOrg := entity.Feed{Name: "team-cymru.org", URL: "https://www.team-cymru.org/Services/Bogons/fullbogons-ipv4.txt", Timeout: 10,
	//	FeedAnalyzers: []entity.FeedAnalyzer{{Score: 3, Expression: "^" + regexpSubnet + ".*"}}}
	//stopforumspamCom := entity.Feed{Name: "stopforumspam.com", URL: "https://www.stopforumspam.com/downloads/toxic_ip_cidr.txt", Timeout: 10,
	//	FeedAnalyzers: []entity.FeedAnalyzer{{Score: 2, Expression: "^" + regexpSubnet + ".*"}}}
	//greensnowCo := entity.Feed{Name: "greensnow.co", URL: "https://blocklist.greensnow.co/greensnow.txt",
	//	Timeout: 10, FeedAnalyzers: []entity.FeedAnalyzer{{Score: 2, Expression: "^" + regexpIP + ".*"}}}
	//binarydefenseCom := entity.Feed{Name: "binarydefense.com", URL: "https://www.binarydefense.com/banlist.txt",
	//	Timeout: 10, FeedAnalyzers: []entity.FeedAnalyzer{{Score: 1, Expression: "^" + regexpIP + ".*"}}}
	//haleysOrgSSH := entity.Feed{Name: "the-haleys.org", URL: "http://charles.the-haleys.org/ssh_dico_attack_with_timestamps.php?days=7",
	//	Timeout: 10, FeedAnalyzers: []entity.FeedAnalyzer{{Score: 1, Expression: "^ALL : ((?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)).*"}}}
	//haleysOrgWp := entity.Feed{Name: "the-haleys.org", URL: "http://charles.the-haleys.org/wp_attack_with_timestamps.php?days=7",
	//	Timeout: 10, FeedAnalyzers: []entity.FeedAnalyzer{{Score: 1, Expression: "^" + regexpIP + ".*"}}}
	//haleysOrgSMTP := entity.Feed{Name: "the-haleys.org", URL: "http://charles.the-haleys.org/smtp_dico_attack_with_timestamps.php?days=7",
	//	Timeout: 10, FeedAnalyzers: []entity.FeedAnalyzer{{Score: 1, Expression: "^" + regexpIP + ".*"}}}
	//blocklistDe := entity.Feed{Name: "blocklist.de", URL: "http://lists.blocklist.de/lists/all.txt", Timeout: 10,
	//	FeedAnalyzers: []entity.FeedAnalyzer{{Score: 3, Expression: "^" + regexpIP + ".*"}}}
	//botscout := entity.Feed{Name: "botscout", URL: "https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/botscout_1d.ipset",
	//	Timeout: 10, FeedAnalyzers: []entity.FeedAnalyzer{{Score: 3, Expression: "^" + regexpIP + ".*"}}}
	//bruteforceblocker := entity.Feed{Name: "bruteforceblocker", URL: "http://danger.rulez.sk/projects/bruteforceblocker/blist.php",
	//	Timeout: 10, FeedAnalyzers: []entity.FeedAnalyzer{{Score: 3, Expression: "^" + regexpIP + ".*"}}}
	//cinsscoreCom := entity.Feed{Name: "cinsscore.com", URL: "http://cinsscore.com/list/ci-badguys.txt", Timeout: 10,
	//	FeedAnalyzers: []entity.FeedAnalyzer{{Score: 3, Expression: "^" + regexpIP + ".*"}}}
	//cruzit := entity.Feed{Name: "cruzit", URL: "https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/cruzit_web_attacks.ipset",
	//	Timeout: 10, FeedAnalyzers: []entity.FeedAnalyzer{{Score: 3, Expression: "^" + regexpIP + ".*"}}}
	//dshieldOrg := entity.Feed{Name: "dshield.org", URL: "http://feeds.dshield.org/top10-2.txt", Timeout: 10,
	//	FeedAnalyzers: []entity.FeedAnalyzer{{Score: 3, Expression: "^" + regexpIP + ".*"}}}
	//emergingthreatsNet := entity.Feed{Name: "emergingthreats.net", URL: "http://rules.emergingthreats.net/open/suricata/rules/compromised-ips.txt",
	//	Timeout: 10, FeedAnalyzers: []entity.FeedAnalyzer{{Score: 3, Expression: "^" + regexpIP + ".*"}}}
	//malwaredomainlist := entity.Feed{Name: "malwaredomainlist", URL: "https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/malwaredomainlist.ipset",
	//	Timeout: 10, FeedAnalyzers: []entity.FeedAnalyzer{{Score: 3, Expression: "^" + regexpIP + ".*"}}}
	//myip := entity.Feed{Name: "myip", URL: "https://myip.ms/files/blacklist/htaccess/latest_blacklist.txt", Timeout: 10,
	//	FeedAnalyzers: []entity.FeedAnalyzer{{Score: 3, Expression: "^deny from ((?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)).*"}}}
	//
	//sslbl := entity.Feed{Name: "sslbl", URL: "https://sslbl.abuse.ch/blacklist/sslipblacklist_aggressive.txt",
	//	Timeout: 10, FeedAnalyzers: []entity.FeedAnalyzer{{Score: 3, Expression: "^" + regexpIP + ".*"}}}
	//rutgersEdu := entity.Feed{Name: "rutgers.edu", URL: "http://report.cs.rutgers.edu/DROP/attackers",
	//	Timeout: 10, FeedAnalyzers: []entity.FeedAnalyzer{{Score: 3, Expression: "^" + regexpIP + ".*"}}}
	//sblamCom := entity.Feed{Name: "sblam.com", URL: "http://sblam.com/blacklist.txt",
	//	Timeout: 10, FeedAnalyzers: []entity.FeedAnalyzer{{Score: 1, Expression: "^" + regexpIP + ".*"}}}
	//talosintelligenceCom := entity.Feed{Name: "talosintelligence.com", URL: "http://www.talosintelligence.com/feeds/ip-filter.blf",
	//	Timeout: 10, FeedAnalyzers: []entity.FeedAnalyzer{{Score: 3, Expression: "^" + regexpIP + ".*"}}}
	//turrisCz := entity.Feed{Name: "turris.cz", URL: "https://www.turris.cz/greylist-data/greylist-latest.csv",
	//	Timeout: 10, FeedAnalyzers: []entity.FeedAnalyzer{{Score: 3, Expression: "^" + regexpIP + ".*"}}}
	//feodotracker := entity.Feed{Name: "feodotracker", URL: "https://feodotracker.abuse.ch/downloads/ipblocklist_aggressive.txt",
	//	Timeout: 10, FeedAnalyzers: []entity.FeedAnalyzer{{Score: 3, Expression: "^" + regexpIP + ".*"}}}
	//fireholdabusers := entity.Feed{Name: "boyscout1d", URL: "https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/firehol_abusers_1d.netset",
	//	Timeout: 10, FeedAnalyzers: []entity.FeedAnalyzer{{Score: 1, Expression: "^" + regexpIP + ".*"}}}
	//var activeFeeds = []entity.Feed{teamCymruOrg, stopforumspamCom, greensnowCo, binarydefenseCom, haleysOrgSSH,
	//	haleysOrgWp, haleysOrgSMTP, spamhaus, firehol, alienvaultCom, badipsCom, blocklistDe, botscout,
	//	bruteforceblocker, cinsscoreCom, cruzit, dshieldOrg, emergingthreatsNet, feodotracker, malwaredomainlist,
	//	myip, sslbl, rutgersEdu, sblamCom, talosintelligenceCom,
	//	turrisCz, feodotracker, fireholdabusers}
	f := entity.Feed{}
	feed, err := f.ReadFile("ip_spam.json")
	if err != nil {
		return nil, nil, err
	}
	ips := make(map[string]entity.IPAnalysis)
	subnets := make(map[string]entity.SUBNETAnalysis)
	for _, activeFeed := range feed {
		c.logger.Printf("[INFO] Importing data feed %s", activeFeed.Name)
		feedResultsIPs, feedResultsSubnets, err := activeFeed.Fetch()
		if err == nil {
			for k, e := range feedResultsIPs { // k is the ip string,  e is the
				if _, ok := ips[k]; ok {
					ip := ips[k]
					ip.Score = ip.Score + e.Score
					ip.Lists = append(ip.Lists, e.Lists[0])
					ips[k] = ip
				} else {
					ips[k] = e
				}
				spam := models.Spam{
					IP:    ips[k].IP,
					Score: ips[k].Score,
				}
				c.iplist = append(c.iplist, spam)

			}
			for k, e := range feedResultsSubnets {
				if _, ok := subnets[k]; ok {
					subnet := subnets[k]
					subnet.Score = subnet.Score + e.Score
					subnet.Lists = append(subnet.Lists, e.Lists[0])
					subnets[k] = subnet
				} else {
					subnets[k] = e
				}
				spam := models.Spam{
					IP:     subnets[k].IP,
					Prefix: subnets[k].PrefixLength,
					Score:  subnets[k].Score,
				}
				c.sublist = append(c.iplist, spam)
			}
			c.logger.Printf("[INFO] Imported %d ips and %d subnets from data feed %s\n", len(feedResultsIPs),
				len(feedResultsSubnets), activeFeed.Name)
		} else {
			c.logger.Printf("[ERROR] Importing data feed %s\n failed : %s", activeFeed.Name, err)
		}
	}

	c.lastUpdate = time.Now()
	return c.iplist, c.sublist, nil

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

func (c *TxtDomains) List() ([]models.Spam, []models.Spam, error) {
	return c.Load(nil)
}
