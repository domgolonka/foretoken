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

This tool saves those threats on multiple different databases and uses REST API to outdata

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
    Date: Fri, 12 Feb 2021 03:21:38 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2

    []

#### Request

`GET /ip/proxy`

    curl -i -H 'Accept: application/json' http://localhost:8080/ip/proxy

#### Response

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Fri, 12 Feb 2021 03:21:38 GMT
    Transfer-Encoding: chunked
    
    {"result":[{"ID":1,"URL":"103.228.xxx.xxx","Type":"ipv4","CreatedAt":"2020-12-04T19:12:05.693099-05:00","UpdatedAt":"2020-12-04T19:12:05.693099-05:00","DeletedAt":null},{"ID":2,"URL":"196.3.xxx.xxx","Type":"ipv4","CreatedAt":"2020-12-04T19:12:05.69557-05:00","UpdatedAt":"2020-12-04T19:12:05.69557-05:00","DeletedAt":null},{"ID":3,"URL":"165.227.xxx.xxx","Type":"ipv4","CreatedAt":"2020-12-04T19:12:05.696224-05:00","UpdatedAt":"2020-12-04T19:12:05.696224-05:00","DeletedAt":null},{"ID":4,"URL":"117.197.xxx.xxx","Type":"ipv4","CreatedAt":"2020-12-04T19:12:05.696876-05:00","UpdatedAt":"2020-12-04T19:12:05.696876-05:00","DeletedAt":null},{"ID":5,"URL":"180.183.xxx.xxx","Type":"ipv4","CreatedAt":"2020-12-04T19:12:05.697515-05:00","UpdatedAt":"2020-12-04T19:12:05.697515-05:00","DeletedAt":null},{"ID":6,"URL":"159.192.xxx.xxx:8080","Type":"ipv4","CreatedAt":"2020-12-04T19:12:05.698074-05:00","UpdatedAt":"2020-12-04T19:12:05.698074-05:00","DeletedAt":null},{"ID":7,"URL":"185.28.xxx.xxx","Type":"ipv4","

#### Request

`GET /ip/spam`

    curl -i -H 'Accept: application/json' http://localhost:8080/ip/spam

#### Response

    HTTP/1.1 200 OK
    Date: Fri, 12 Feb 2021 03:29:54 GMT
    Content-Type: text/plain; charset=utf-8
    Transfer-Encoding: chunked
    
    205.159.xxx.xxx/24
    103.62.xxx.xxx/22
    103.245.xxx.xxx/23
    209.161.xxx.xxx/19

#### Request

`GET /ip/vpn`

    curl -i -H 'Accept: application/json' http://localhost:8080/ip/vpn

#### Response

    HTTP/1.1 200 OK
    Date: Fri, 12 Feb 2021 03:29:54 GMT
    Content-Type: text/plain; charset=utf-8
    Transfer-Encoding: chunked

    yul-c14.xxx.com
    lim-c04.xxx.com
    bhx-c05.xxx.com

#### Request

`GET /ip/tor`

    curl -i -H 'Accept: application/json' http://localhost:8080/ip/tor

#### Response

    HTTP/1.1 200 OK
    Date: Fri, 12 Feb 2021 03:29:54 GMT
    Content-Type: text/plain; charset=utf-8
    Transfer-Encoding: chunked

    205.159.xxx.xxx/24
    103.62.xxx.xxx/22
    103.245.xxx.xxx/23
    209.161.xxx.xxx/19

#### Request

`GET /email/disposal`

    curl -i -H 'Accept: application/json' http://localhost:8080/email/disposal

#### Response

    HTTP/1.1 200 OK
    Date: Fri, 12 Feb 2021 03:29:54 GMT
    Content-Type: text/plain; charset=utf-8
    Transfer-Encoding: chunked

    xxx.cc
    xxx.com
    xxx.ca


#### Request

`GET /email/generic`

    curl -i -H 'Accept: application/json' http://localhost:8080/email/generic

#### Response

    HTTP/1.1 200 OK
    Date: Fri, 12 Feb 2021 03:29:54 GMT
    Content-Type: text/plain; charset=utf-8
    Transfer-Encoding: chunked

    xxx.cc
    xxx.com
    xxx.ca

#### Request

`GET /email/spam`

    curl -i -H 'Accept: application/json' http://localhost:8080/email/spam

#### Response

    HTTP/1.1 200 OK
    Date: Fri, 12 Feb 2021 03:29:54 GMT
    Content-Type: text/plain; charset=utf-8
    Transfer-Encoding: chunked

    xxx.cc
    xxx.com
    xxx.ca

#### Request

`GET /score/email?email=youremail@gmail.com`

    curl -i -H 'Accept: application/json' http://localhost:8080/score/email?email=youremail@gmail.com

#### Response

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