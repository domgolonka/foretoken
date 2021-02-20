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
	Type  string
	Lists []string // list of websites where it was found
}

type SUBNETAnalysis struct {
	IP           string
	SUBNET       string
	Type         string
	PrefixLength uint8
	Score        int
	Lists        []string // list of websites where it was found
}

type FeedAnalyzer struct {
	Score      int    `json:"score"`
	Expression string `json:"expression"`
}

type Expression struct {
	Name       string `json:"name"`
	Expression string `json:"expression"`
	Type       string `json:"type"`
}

type Feed struct {
	Name          string         `json:"name"`
	URL           string         `json:"url"`
	Type          string         `json:"type"`
	Format        string         `json:"format"`
	Timeout       time.Duration  `json:"timeout"`
	FeedAnalyzers []FeedAnalyzer `json:"feed"`
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

func (feed Feed) GetExpressions() ([]string, error) {
	var expressions []string
	workingdir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	var f []*Expression
	file, err := ioutil.ReadFile(workingdir + "/resource/expressions.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &f)
	if err != nil {
		return nil, err
	}

	for _, s := range feed.FeedAnalyzers {
		for _, a := range f {
			if strings.EqualFold(s.Expression, a.Name) {
				expressions = append(expressions, a.Expression)
			}
		}

	}
	return expressions, nil
}

func (feed Feed) FetchString() (map[string]string, error) {
	var netClient = &http.Client{
		Timeout: time.Second * feed.Timeout,
	}
	response, err := netClient.Get(feed.URL)

	if err != nil {
		return nil, err
	}

	defer func() {
		err = response.Body.Close()
	}()

	workingdir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	var expfile []*Expression
	file, err := ioutil.ReadFile(workingdir + "/resource/expressions.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &expfile)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(response.Body)
	scanner.Split(bufio.ScanRunes)
	var buf bytes.Buffer
	for scanner.Scan() {
		buf.WriteString(scanner.Text())
	}
	var httpResult = buf.String()
	resultString := make(map[string]string)
	for _, element := range strings.Split(httpResult, "\n") {
		line := strings.Trim(element, " ")

		match := false
		for _, fa := range feed.FeedAnalyzers {
			for _, a := range expfile {
				if strings.EqualFold(fa.Expression, a.Name) {

					regex, _ := regexp.Compile(`` + a.Expression + ``)
					var findings = regex.FindStringSubmatch(line)
					if !match {
						if len(findings) == 2 {
							resultString[findings[1]] = findings[1]
							match = true
						}
					}
				}
			}
		}
	}
	return resultString, nil
}

func (feed Feed) FetchIP() (map[string]IPAnalysis, map[string]SUBNETAnalysis, error) {
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

	workingdir, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}
	var expfile []*Expression
	file, err := ioutil.ReadFile(workingdir + "/resource/expressions.json")
	if err != nil {
		return nil, nil, err
	}
	err = json.Unmarshal(file, &expfile)
	if err != nil {
		return nil, nil, err
	}

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
			for _, a := range expfile {
				if strings.EqualFold(fa.Expression, a.Name) {

					regex, _ := regexp.Compile(`` + a.Expression + ``)
					var findings = regex.FindStringSubmatch(line)
					if !match {
						if len(findings) == 2 {
							resultIps[findings[1]] = IPAnalysis{findings[1], fa.Score, a.Type, []string{feed.Name}}
							match = true
						} else if len(findings) == 3 {
							subnet := findings[1] + "/" + findings[2]
							prefixLength, _ := strconv.Atoi(findings[2])
							resultSubnets[subnet] = SUBNETAnalysis{findings[1], subnet, a.Type, uint8(prefixLength),
								fa.Score, []string{feed.Name}}
							match = true
						}
					}
				}
			}
		}
	}
	return resultIps, resultSubnets, nil
}
