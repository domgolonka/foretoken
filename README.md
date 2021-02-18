<p align="center"> 
  <img src="assets/img.png" width="300" height="300" alt="Threat Defender" /></p>
  <h2 align="center">Threat Defender</h2>
  <p align="center">A modern-day defence tool using REST/gRPC protocols.</p>

<p align="center">
    <a href="https://github.com/domgolonka/threatdefender/issues/new/choose">Report Bug</a>
    ·
    <a href="https://github.com/domgolonka/threatdefender/issues/new/choose">Request Feature</a>
</p>
<p align="center">Loved the project? Please consider donating to the bitcoin address below:</p>

<p align="center"><img src="assets/bitcoinaddress.png" alt="3Gn3URoFijqx2keY1fAfSpf8kZge5MBDGM" height="300" width="300" border="0" /> </p>
<p align="center">Note: This is still in alpha release, this might break over time. </p>

## About

ThreatDefender is a tool to scrape for potential dangerous threats faced on the internet. The list of threats scraped at
the moment is:

- Emails
    - Disposable
    - Generic
    - Free
    - Spam
- IPs
    - VPN
    - Spam
    - Proxy
    - Tor
  
- Score
    - IP [(0 to 100)](#score)
    - Email [(0 to 100)](#score)

### Score

The overall Fraud Score of the email and IP's reputation and recent behavior across the threat network. Fraud Scores >= 75 are suspicious, but not necessarily fraudulent.


This tool saves those threats on multiple different databases and uses REST API & gRPC to output data.

# Usage

## Migrate

**BEFORE YOU RUN THIS**, You need to migrate the database:

`make migrate`

## How to run

To run it on your local computer:

`make run`

The default config file is `config/config.dev.yml`. If you want to run it with a different config file (or add your own).

`make build` (make sure to build it first)

`./bin/threatdefender --config=/PATH/TO/CONFIG`

example:
`./bin/threatdefender --config=config/config.prod.yml`

# Configs

All configurations are in the config files in the directory "config". You can add your own environment variables here.

## Change the databases

At this moment, Threat Defender only supports SQLite and PostgreSQL. You can change the `databasename` field with either `postgresql` or `sqlite3`

For Postgresql, I would advise using a quick read/write database like [timescale](https://www.timescale.com/). 

# APIs

REST API & gRPC is enabled.

### gRPC 

The default gRPC port is 8082 (you can change in the config)


### REST API

The REST API to the example app is described below.

#### Request

`GET /health`

    curl -i -H 'Accept: application/json' http://localhost:8080/health

#### Response

    HTTP/1.1 200 OK
    Date: Thu, 18 Feb 2021 04:56:45 GMT
    Content-Type: application/json
    Content-Length: 13
    
    {"http":true}

#### Request

`GET /list/ip/proxy`

    curl -i -H 'Accept: application/json' http://localhost:8080/list/ip/proxy

#### Response

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Fri, 12 Feb 2021 03:21:38 GMT
    Transfer-Encoding: chunked
    
    {"result":[{"ID":1,"URL":"103.228.xxx.xxx","Type":"ipv4","CreatedAt":"2020-12-04T19:12:05.693099-05:00","UpdatedAt":"2020-12-04T19:12:05.693099-05:00","DeletedAt":null},{"ID":2,"URL":"196.3.xxx.xxx","Type":"ipv4","CreatedAt":"2020-12-04T19:12:05.69557-05:00","UpdatedAt":"2020-12-04T19:12:05.69557-05:00","DeletedAt":null},{"ID":3,"URL":"165.227.xxx.xxx","Type":"ipv4","CreatedAt":"2020-12-04T19:12:05.696224-05:00","UpdatedAt":"2020-12-04T19:12:05.696224-05:00","DeletedAt":null},{"ID":4,"URL":"117.197.xxx.xxx","Type":"ipv4","CreatedAt":"2020-12-04T19:12:05.696876-05:00","UpdatedAt":"2020-12-04T19:12:05.696876-05:00","DeletedAt":null},{"ID":5,"URL":"180.183.xxx.xxx","Type":"ipv4","CreatedAt":"2020-12-04T19:12:05.697515-05:00","UpdatedAt":"2020-12-04T19:12:05.697515-05:00","DeletedAt":null},{"ID":6,"URL":"159.192.xxx.xxx:8080","Type":"ipv4","CreatedAt":"2020-12-04T19:12:05.698074-05:00","UpdatedAt":"2020-12-04T19:12:05.698074-05:00","DeletedAt":null},{"ID":7,"URL":"185.28.xxx.xxx","Type":"ipv4","

#### Request

`GET /list/ip/spam`

    curl -i -H 'Accept: application/json' http://localhost:8080/list/ip/spam

#### Response

    HTTP/1.1 200 OK
    Date: Thu, 18 Feb 2021 04:57:33 GMT
    Content-Type: text/plain; charset=utf-8
    Content-Length: 34952
    
    168.0.xxx.0/22
    202.49.xxx.0/24

#### Request

`GET /list/ip/vpn`

    curl -i -H 'Accept: application/json' http://localhost:8080/list/ip/vpn

#### Response

    HTTP/1.1 200 OK
    Date: Fri, 12 Feb 2021 03:29:54 GMT
    Content-Type: text/plain; charset=utf-8
    Transfer-Encoding: chunked

    yul-c14.xxx.com
    lim-c04.xxx.com
    bhx-c05.xxx.com

#### Request

`GET /list/ip/tor`

    curl -i -H 'Accept: application/json' http://localhost:8080/list/ip/tor

#### Response

    HTTP/1.1 200 OK
    Date: Thu, 18 Feb 2021 04:58:18 GMT
    Content-Type: text/plain; charset=utf-8
    Content-Length: 23253
    
    176.10.xxx.xxx
    54.37.xxx.xxx
    109.70.xxx.xxx

#### Request

`GET /list/email/disposal`

    curl -i -H 'Accept: application/json' http://localhost:8080/list/email/disposal

#### Response

    HTTP/1.1 200 OK
    Date: Thu, 18 Feb 2021 04:58:18 GMT
    Content-Type: text/plain; charset=utf-8
    Content-Length: xxx

    xxx.cc
    xxx.com
    xxx.ca


#### Request

`GET /list/email/generic`

    curl -i -H 'Accept: application/json' http://localhost:8080/list/email/generic

#### Response

    HTTP/1.1 200 OK
    Date: Thu, 18 Feb 2021 04:59:38 GMT
    Content-Type: text/plain; charset=utf-8
    Content-Length: xxxx
    
    xxx@
    xxx@
    xxx@

#### Request

`GET /list/email/spam`

    curl -i -H 'Accept: application/json' http://localhost:8080/list/email/spam

#### Response

    HTTP/1.1 200 OK
    Date: Thu, 18 Feb 2021 04:59:38 GMT
    Content-Type: text/plain; charset=utf-8
    Content-Length: xxxx

    xxx.cc
    xxx.com
    xxx.ca

#### Request

`GET /list/email/free`

    curl -i -H 'Accept: application/json' http://localhost:8080/list/email/free

#### Response

    HTTP/1.1 200 OK
    Date: Thu, 18 Feb 2021 04:59:38 GMT
    Content-Type: text/plain; charset=utf-8
    Content-Length: xxxx

    xxx.cc
    xxx.com
    xxx.ca

#### Request

`GET /score/email/youremail@gmail.com`

    curl -i -H 'Accept: application/json' http://localhost:8080/score/email/youremail@gmail.com

#### Response

    HTTP/1.1 200 OK
    Date: Fri, 12 Feb 2021 03:29:54 GMT
    Content-Type: text/plain; charset=utf-8
    Transfer-Encoding: chunked

    10

#### Request

`GET /score/ip/127.0.0.1`

    curl -i -H 'Accept: application/json' http://localhost:8080/score/ip/127.0.0.1

#### Response

    HTTP/1.1 200 OK
    Date: Fri, 12 Feb 2021 03:29:54 GMT
    Content-Type: text/plain; charset=utf-8
    Transfer-Encoding: chunked

    0


#### Request

`GET /validate/email/youremail@gmail.com`

    curl -i -H 'Accept: application/json' http://localhost:8080/validate/email/youremail@gmail.com

#### Response

    HTTP/1.1 200 OK
    Date: Fri, 12 Feb 2021 03:29:54 GMT
    Content-Type: text/json; charset=utf-8
    Transfer-Encoding: chunked

    {
    "valid": true
    }

#### Request

`GET /email/youremail@gmail.com`

    curl -i -H 'Accept: application/json' http://localhost:8080/email/youremail@gmail.com

#### Response

    HTTP/1.1 200 OK
    Date: Fri, 12 Feb 2021 03:29:54 GMT
    Content-Type: text/json; charset=utf-8
    Transfer-Encoding: chunked

    {
    "valid": true,
    "disposable": false,
    "recent_spam": false,
    "free": false,
    "leaked": false,
    "generic": false,
    "score": 0,
    "domain": {
        "created_at": "1995-08-13T04:00:00Z",
        "expiration_date": "2021-08-12T04:00:00Z"
        }
    }

#### Request

`GET /ip/127.0.0.1`

    curl -i -H 'Accept: application/json' http://localhost:8080/ip/127.0.0.1

#### Response

    HTTP/1.1 200 OK
    Date: Fri, 12 Feb 2021 03:29:54 GMT
    Content-Type: text/json; charset=utf-8
    Transfer-Encoding: chunked

    {
    "success": false,
    "proxy": false,
    "ISP": "",
    "organization": "",
    "ASN": 0,
    "host": "",
    "country_code": "",
    "city": "",
    "region": "",
    "is_crawler": false,
    "connection_type": "",
    "latitude": 0,
    "longitude": 0,
    "timezone": "",
    "vpn": false,
    "tor": false,
    "recent_abuse": false,
    "abuse_velocity": "",
    "bot_status": false,
    "mobile": false,
    "score": 0,
    "operating_system": "",
    "browser": "",
    "device_model": "",
    "device_brand": ""
    }




## Work in progress

Lots of features are being worked on.

## Roadmap

I would like a discussion going on the potential expansion of the tool.

I would like this tool to detect all modern threats.