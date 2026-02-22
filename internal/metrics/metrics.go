package metrics

import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	SVIDIssued = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "zt_svid_issued_total",
		Help: "Total number of SVIDs issued",
	})
	SVIDRevoked = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "zt_svid_revoked_total",
		Help: "Total number of SVIDs revoked",
	})
	AuthFailures = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "zt_auth_failures_total",
		Help: "Total number of authentication failures",
	})
)

func init() {
	prometheus.MustRegister(SVIDIssued)
	prometheus.MustRegister(SVIDRevoked)
	prometheus.MustRegister(AuthFailures)
}

func Handler() http.Handler {
	return promhttp.Handler()
}
