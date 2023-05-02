module github.com/domgolonka/foretoken

go 1.17

// replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible

// replace github.com/coreos/bbolt v1.3.5 => go.etcd.io/bbolt v1.3.5

require (
	github.com/Boostport/address v0.6.0
	github.com/airbrake/gobrake v3.7.4+incompatible
	github.com/ansrivas/fiberprometheus/v2 v2.1.1
	github.com/antchfx/htmlquery v1.2.3
	github.com/araddon/dateparse v0.0.0-20210207001429-0eec95c9db7e
	github.com/caio/go-tdigest v3.1.0+incompatible // indirect
	github.com/domainr/whois v0.0.0-20210215110205-c05ecdd18962
	github.com/getsentry/sentry-go v0.9.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/gofiber/fiber/v2 v2.29.0
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/schema v1.2.0
	github.com/hashicorp/consul/api v1.3.0
	github.com/hashicorp/go-multierror v1.1.0
	github.com/jinzhu/configor v1.2.1
	github.com/jmoiron/sqlx v1.3.1
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.9.0
	github.com/likexian/whois-parser-go v1.15.2
	github.com/mattn/go-sqlite3 v1.14.6
	github.com/oschwald/geoip2-golang v1.6.1
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/samuel/go-zookeeper v0.0.0-20190923202752-2cc03de413da
	github.com/sirupsen/logrus v1.7.0
	github.com/soluchok/freeproxy v0.0.0-20200112224202-ccb33291a087
	github.com/soluchok/go-cloudflare-scraper v0.0.0-20190117212330-ecf651d4e614
	github.com/test-go/testify v1.1.4
	golang.org/x/crypto v0.1.0
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
)

require github.com/common-nighthawk/go-figure v0.0.0-20210622060536-734e95fb86be

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/PuerkitoBio/goquery v1.6.1 // indirect
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/andybalholm/cascadia v1.2.0 // indirect
	github.com/antchfx/xpath v1.1.6 // indirect
	github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/creack/pty v1.1.11 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fatih/color v1.9.0 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/gofiber/adaptor/v2 v2.1.1 // indirect
	github.com/gofiber/utils v0.1.2 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.1 // indirect
	github.com/hashicorp/go-immutable-radix v1.0.0 // indirect
	github.com/hashicorp/go-rootcerts v1.0.0 // indirect
	github.com/hashicorp/golang-lru v0.5.1 // indirect
	github.com/hashicorp/serf v0.8.2 // indirect
	github.com/klauspost/compress v1.15.0 // indirect
	github.com/leesper/go_rng v0.0.0-20190531154944-a612b043e353 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/likexian/gokit v0.23.3 // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/oschwald/maxminddb-golang v1.8.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.9.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.15.0 // indirect
	github.com/prometheus/procfs v0.2.0 // indirect
	github.com/robertkrimen/otto v0.0.0-20180617131154-15f95af6e78d // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/soluchok/gokogiri v0.0.0-20190903214353-0718098bc8db // indirect
	github.com/unchartedsoftware/witch v0.0.0-20200617171400-4f405404126f // indirect
	github.com/urfave/cli v1.22.4 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.34.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	github.com/zonedb/zonedb v1.0.3021 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.1.0 // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	gonum.org/v1/gonum v0.9.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)
