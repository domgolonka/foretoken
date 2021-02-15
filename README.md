# ThreatScraper

Note: This is still in alpha release, this might break over time.

## About
![logo](assets/img.png){height="36px" width="36px"}


ThreatScraper is a tool to scrape for potential dangerous threats faced on the internet. The list of threats scraped at
the moment is:

- Emails
    - Disposable
- IPs
    - VPN
    - Spam
    - Proxy
    - Tor

This tool saves those threats on multiple different databases and uses REST API to outdata

# Usage

## Migrate

**BEFORE YOU RUN THIS**, You need to migrate the database:

`make migrate`

## How to run

To run it on your local computer:

`make run`

## Changing ports and configs

All configurations are in the config files in the directory "config". You can add your own environment variables here.

# Apis

## REST API

The REST API to the example app is described below.

### Request

`GET /public/health`

    curl -i -H 'Accept: application/json' http://localhost:8080/public/health

### Response

    HTTP/1.1 200 OK
    Date: Fri, 12 Feb 2021 03:21:38 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2

    []

### Request

`GET /public/ip/proxy`

    curl -i -H 'Accept: application/json' http://localhost:8080/public/ip/proxy

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Fri, 12 Feb 2021 03:21:38 GMT
    Transfer-Encoding: chunked
    
    {"result":[{"ID":1,"URL":"103.228.xxx.xxx","Type":"ipv4","CreatedAt":"2020-12-04T19:12:05.693099-05:00","UpdatedAt":"2020-12-04T19:12:05.693099-05:00","DeletedAt":null},{"ID":2,"URL":"196.3.xxx.xxx","Type":"ipv4","CreatedAt":"2020-12-04T19:12:05.69557-05:00","UpdatedAt":"2020-12-04T19:12:05.69557-05:00","DeletedAt":null},{"ID":3,"URL":"165.227.xxx.xxx","Type":"ipv4","CreatedAt":"2020-12-04T19:12:05.696224-05:00","UpdatedAt":"2020-12-04T19:12:05.696224-05:00","DeletedAt":null},{"ID":4,"URL":"117.197.xxx.xxx","Type":"ipv4","CreatedAt":"2020-12-04T19:12:05.696876-05:00","UpdatedAt":"2020-12-04T19:12:05.696876-05:00","DeletedAt":null},{"ID":5,"URL":"180.183.xxx.xxx","Type":"ipv4","CreatedAt":"2020-12-04T19:12:05.697515-05:00","UpdatedAt":"2020-12-04T19:12:05.697515-05:00","DeletedAt":null},{"ID":6,"URL":"159.192.xxx.xxx:8080","Type":"ipv4","CreatedAt":"2020-12-04T19:12:05.698074-05:00","UpdatedAt":"2020-12-04T19:12:05.698074-05:00","DeletedAt":null},{"ID":7,"URL":"185.28.xxx.xxx","Type":"ipv4","

### Request

`GET /public/ip/spam`

    curl -i -H 'Accept: application/json' http://localhost:8080/public/ip/spam

### Response

    HTTP/1.1 200 OK
    Date: Fri, 12 Feb 2021 03:29:54 GMT
    Content-Type: text/plain; charset=utf-8
    Transfer-Encoding: chunked
    
    205.159.xxx.xxx/24
    103.62.xxx.xxx/22
    103.245.xxx.xxx/23
    209.161.xxx.xxx/19

### Request

`GET /public/ip/vpn`

    curl -i -H 'Accept: application/json' http://localhost:8080/public/ip/vpn

### Response

    HTTP/1.1 200 OK
    Date: Fri, 12 Feb 2021 03:29:54 GMT
    Content-Type: text/plain; charset=utf-8
    Transfer-Encoding: chunked

    yul-c14.xxx.com
    lim-c04.xxx.com
    bhx-c05.xxx.com

### Request

`GET /public/ip/tor`

    curl -i -H 'Accept: application/json' http://localhost:8080/public/ip/tor

### Response

    HTTP/1.1 200 OK
    Date: Fri, 12 Feb 2021 03:29:54 GMT
    Content-Type: text/plain; charset=utf-8
    Transfer-Encoding: chunked

    205.159.xxx.xxx/24
    103.62.xxx.xxx/22
    103.245.xxx.xxx/23
    209.161.xxx.xxx/19

### Request

`GET /public/email/disposal`

    curl -i -H 'Accept: application/json' http://localhost:8080/public/email/disposal

### Response

    HTTP/1.1 200 OK
    Date: Fri, 12 Feb 2021 03:29:54 GMT
    Content-Type: text/plain; charset=utf-8
    Transfer-Encoding: chunked

    xxx.cc
    xxx.com
    xxx.ca

## Work in progress

Lots of features are being worked on.

## Roadmap

I would like a discussion going on the potential expansion of the tool.

I would like this tool to detect all modern threats.