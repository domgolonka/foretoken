package metrics

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
)

//var (
//	httpRequests = prometheus.NewCounterVec(
//		prometheus.CounterOpts{
//			Name: "http_requests_total",
//			Help: "How many HTTP requests processed, partitioned by name and status code",
//		},
//		[]string{"name", "code"},
//	)
//	httpTimings = prometheus.NewHistogramVec(
//		prometheus.HistogramOpts{
//			Name:    "http_request_times",
//			Help:    "The duration of HTTP requests, partitioned by name",
//			Buckets: prometheus.ExponentialBuckets(0.001, 10, 5),
//		},
//		[]string{"name"},
//	)
//)
//
//func init() {
//	prometheus.MustRegister(httpRequests)
//	prometheus.MustRegister(httpTimings)
//}

func InitPrometheus(app *fiber.App) fiber.Router {
	// middleware
	prom := fiberprometheus.New("threatdefender")
	prom.RegisterAt(app, "/metrics")
	return app.Use(prom.Middleware)
}

/*func InstrumentRoute(name string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metrics := httpsnoop.CaptureMetrics(next, w, r)
		httpRequests.WithLabelValues(name, strconv.Itoa(metrics.Code)).Inc()
		httpTimings.WithLabelValues(name).Observe(float64(metrics.Duration.Seconds()))
	})
}*/
