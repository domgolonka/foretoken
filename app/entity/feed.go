package entity

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type IPAnalysis struct {
	IP    string
	Score int
	Lists []string
}

type SUBNETAnalysis struct {
	SUBNET       string
	PrefixLength byte
	Score        int
	Lists        []string
}

type FeedAnalyzer struct {
	Score      int
	Expression string
}

type Feed struct {
	Name          string
	URL           string
	Type          string
	Format        string
	Timeout       time.Duration
	FeedAnalyzers []FeedAnalyzer
}

func (feed Feed) ReadFile(filename string) ([]*Feed, error) {
	var f []*Feed
	workingdir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	file, err := ioutil.ReadFile(workingdir + "/resource/" + filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &f)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (feed Feed) Fetch() (map[string]IPAnalysis, map[string]SUBNETAnalysis, error) {
	var netClient = &http.Client{
		Timeout: time.Second * feed.Timeout,
	}
	response, err := netClient.Get(feed.URL)

	if err != nil {
		return nil, nil, err
	}

	defer func() {
		err = response.Body.Close()
	}()

	scanner := bufio.NewScanner(response.Body)
	scanner.Split(bufio.ScanRunes)
	var buf bytes.Buffer
	for scanner.Scan() {
		buf.WriteString(scanner.Text())
	}
	var httpResult = buf.String()

	resultIps := make(map[string]IPAnalysis)
	resultSubnets := make(map[string]SUBNETAnalysis)
	for _, element := range strings.Split(httpResult, "\n") {
		line := strings.Trim(element, " ")

		match := false
		for _, fa := range feed.FeedAnalyzers {
			regex, _ := regexp.Compile(fa.Expression)
			var findings = regex.FindStringSubmatch(line)
			if !match {
				if len(findings) == 2 {
					resultIps[findings[1]] = IPAnalysis{findings[1], fa.Score, []string{feed.Name}}
					match = true
				} else if len(findings) == 3 {
					subnet := findings[1] + "/" + findings[2]
					prefixLength, _ := strconv.Atoi(findings[2])
					resultSubnets[subnet] = SUBNETAnalysis{subnet, byte(prefixLength),
						fa.Score, []string{feed.Name}}
					match = true
				}
			}
		}
	}
	return resultIps, resultSubnets, nil
}
