<p align="center"> 
  <img src="assets/img.png" width="300" height="300" alt="Foretoken" /></p>
  <h2 align="center">Foretoken</h2>
  <p align="center">A blazing fast, highly customizable, modern-day defence tool using (in memory) SQL & REST/gRPC protocols.</p>

<p align="center">
    <a href="https://github.com/domgolonka/foretoken/issues/new/choose">Report Bug</a>
    ·
    <a href="https://github.com/domgolonka/foretoken/issues/new/choose">Request Feature</a>
</p>

<p align="center"><a href="https://foretoken.domgolonka.com">https://foretoken.domgolonka.com</a></p>

<p align="center">Loved the project? Please consider donating to the bitcoin address below:</p>

<p align="center"><img src="assets/bitcoinaddress.png" alt="3Gn3URoFijqx2keY1fAfSpf8kZge5MBDGM" height="300" width="300" border="0" /> </p>


## About

Foretoken is a tool to scrape and defend against potential dangerous threats faced on the internet. It aims to be a
highly customizable tools for companies and individuals to use to counter threats.

### Features

- **Emails** (Rest/gRPC)
    - Disposable
    - Generic
    - Free
    - Spam
- **IPs** (Rest/gRPC)
    - VPN
    - Spam
    - Proxy
    - Tor
- **Score** (Rest/gRPC)
    - IP [(0 to 100)](#score)
    - Email [(0 to 100)](#score)
- **Database**
    - SQLite
    - PostgreSQL
- **[Editable Sources](#source)**
    - You can edit all sources

# Usage

## Migrate

**If using NON-MEMORY SQLITE or PostgreSQL, DO THIS BEFORE YOU RUN**, You need to migrate the database:

`make migrate`

## How to run

To run it on your local computer:

    git clone https://github.com/domgolonka/foretoken
    cd ./foretoken
    make build && ./bin/foretoken

The default config file is `config.yml`. 
If you want to run it with a different config file (or add your own).


`git clone https://github.com/domgolonka/foretoken`
`make build` (make sure to build it first)

`./bin/foretoken --config=/PATH/TO/CONFIG`

example:
`./bin/foretoken --config=./config.prod.yml`

## Docker

You can run it in docker, locally:

```docker build -t foretoken .```

Once the image is built, Foretoken can be invoked by running the following:

```docker run --rm -t -p 8080:8080 foretoken ```

Or run Docker from our repo:

    docker run -d -p 8080:8080 domgolonka/foretoken

or with a custom config file:

```docker run -d -p 8080:8080 domgolonka/foretoken --config=config.yml```

# Configs

All configurations are in the config files in the directory "config". You can add your own environment variables here.

## External APIs

The application is improved if you sign up for external APIs. Leaked is paid, but all other services are free to sign up!

- [haveibeenpwned.com](https://haveibeenpwned.com/) - Check if email/password is leaked.
- [maxmind.com](https://www.maxmind.com/en/home) - IP Geolocation
- [ip2location.com](https://www.ip2location.com/) - IP Geolocation

Change the file `changeme.env` to `.env` and save any External API Keys.


    PWNEDKEY=
    IP2LOCATION=
    MAXMIND=

For full configuration examples, check out [https://foretoken.domgolonka.com](https://foretoken.domgolonka.com)

## Change the databases

At this moment, Foretoken only supports SQLite and PostgreSQL. You can change the `databasename` field with
either `postgresql` or `sqlite3`

By Default, the SQLite driver is set to "in memory". To use a file, you need to
specify that the `host` to a `.sqlite3` extension, example: `YOURNAME.sqlite3`. This will create a
new SQLite file  in the root directory.

For Postgresql, I would advise using a quick read/write database like [timescale](https://www.timescale.com/).

***PostgreSQL is not yet tested***

## Score

The overall Fraud Score of the email and IP's reputation and recent behavior across the threat network. Fraud Scores >=
75 are suspicious, but not necessarily fraudulent.

This tool saves those threats on multiple different databases and uses REST API & gRPC to output data.

## Source

All sources are available in the `./resource` directory. You can edit and the resources. They files get checked once a
day by the different modules.

### Regular Expressions

Regex expressions are saved in the `./resource/expressions.json` file in JSON format.

Each regex looks like this:

    {
    "name": "ipv4",
    "expression": "^((?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)).*",
    "type": "ipv4"
    },

_**Name**:_  The UNIQUE name of the regular expression. 

_**Expression**:_ The regex 

_**Type:**_ The type of expression. For IPs, it is
usually is a ipv4 or ipv6. For IP proxy, its http, https, sock4, sock5.


The files are stored in the `./resource` directory and start with `ip_` such as `ip_tor` for tor.

You can add sources by adding a new file to `./resource` directory and updating the `config.yml` file:

    ### Resource files
    resource:
      emaildisposallist: [ "email_disposable" ]
      emailfreelist: [ "email_free" ]
      emailspamlist: [ "email_spam" ]
      ipvpnlist: [ "ip_vpn" ]
      ipopenvpnlist: [ "ip_openvpn" ]
      iptorlist: [ "ip_tor" ]
      ipproxylist: [ "ip_proxy" ]
      ipspamlist: [ "ip_spam" ]
      expressionlist: [ "expressions" ]


# APIs

REST API & gRPC is enabled. For more API examples: [https://foretoken.domgolonka.com](https://foretoken.domgolonka.com)

## gRPC

The default gRPC port is 8082 (you can change in the config)

## REST API

The REST API to the example app is described below.

### Rate Limiting

You can enable the rate limiter for REST API in the `config.yml` file.

    ratelimit:
      enabled: true
      max: 20 
      expiration: 30 

`Max` number of recent connections during `Duration` seconds before sending a 429 response
 
`Expiration` is the time on how long to keep records of requests in memory per minute

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

# Service Discovery

Foretoken supports etcd3, zookeeper, and consul as a registry.

All service discovery configurations are stored in the `config.yml` file:

    servicediscovery:
      service: ""
      nodeid: ""
      endpoint: ""

- Service: The viable options are `consul`, `etc3` and `zookeeper`
- Nodeid: A name for the grpc nodeid
- endpoint: An address for the service such as zookeeper: `10.0.101.68:2189`, etcd: `http://10.0.101.68:2379`  or consul: `http://10.0.101.68:8500`

# Metrics

## Prometheus

Prometheus is enabled. Following metrices are available by default:

    http_requests_total
    http_request_duration_seconds
    http_requests_in_progress_total

## Work in progress

Lots of features are being worked on.

## Roadmap

I would like a discussion going on the potential expansion of the tool.

I would like this tool to detect all modern threats.