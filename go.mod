module github.com/domgolonka/foretoken

go 1.15

replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible

replace github.com/coreos/bbolt v1.3.5 => go.etcd.io/bbolt v1.3.5

require (
	github.com/Boostport/address v0.6.0
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/airbrake/gobrake v3.7.4+incompatible
	github.com/ansrivas/fiberprometheus/v2 v2.1.1
	github.com/antchfx/htmlquery v1.2.3
	github.com/araddon/dateparse v0.0.0-20210207001429-0eec95c9db7e
	github.com/caio/go-tdigest v3.1.0+incompatible // indirect
	github.com/common-nighthawk/go-figure v0.0.0-20200609044655-c4b36f998cf2
	github.com/coreos/bbolt v1.3.5 // indirect
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/domainr/whois v0.0.0-20210215110205-c05ecdd18962
	github.com/etcd-io/etcd v3.3.25+incompatible
	github.com/getsentry/sentry-go v0.9.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/gofiber/fiber/v2 v2.5.0
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/schema v1.2.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/hashicorp/consul/api v1.3.0
	github.com/hashicorp/go-multierror v1.1.0
	github.com/jinzhu/configor v1.2.1
	github.com/jmoiron/sqlx v1.3.1
	github.com/joho/godotenv v1.3.0
	github.com/kr/text v0.2.0 // indirect
	github.com/leesper/go_rng v0.0.0-20190531154944-a612b043e353 // indirect
	github.com/lib/pq v1.9.0
	github.com/likexian/whois-parser-go v1.15.2
	github.com/mattn/go-sqlite3 v1.14.6
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/oschwald/geoip2-golang v1.5.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/samuel/go-zookeeper v0.0.0-20190923202752-2cc03de413da
	github.com/sirupsen/logrus v1.7.0
	github.com/soluchok/freeproxy v0.0.0-20200112224202-ccb33291a087
	github.com/soluchok/go-cloudflare-scraper v0.0.0-20190117212330-ecf651d4e614
	github.com/test-go/testify v1.1.4
	golang.org/x/crypto v0.0.0-20200707235045-ab33eee955e0
	golang.org/x/exp v0.0.0-20191129062945-2f5052295587 // indirect
	golang.org/x/tools v0.1.0 // indirect
	gonum.org/v1/gonum v0.8.2 // indirect
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
)
