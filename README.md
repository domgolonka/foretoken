# ThreatScraper

Note: This is still in alpha release, this might break over time.


## About
ThreatScraper is a tool to scrape for potential dangerous threats faced on the internet. 
The list of threats scraped at the moment is:
- Emails
  - Disposable
- IPs
  - VPN 
  - Spam
  - Proxy
    

This tool saves those threats on multiple different databases and uses REST API to outdata

# Usage 
## How to run

To run it on your local computer:

`make run`

## Changing ports and configs

Everything is in the Config file app/config. You can add your own environment variables here.

# Apis

## REST API

The REST API to the example app is described below.


### Request

`GET /public/health`

    curl -i -H 'Accept: application/json' http://localhost:8080/public/health

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2

    []

### Request

`GET /public/proxy`

    curl -i -H 'Accept: application/json' http://localhost:8080/public/ip/proxy

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2

    []


## Work in progress

Lots of features are being worked on.

## Roadmap
I would like a discussion going on the potential expansion of the tool.

I would like this tool to detect all modern threats with machine learn.