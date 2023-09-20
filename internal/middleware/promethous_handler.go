package middleware

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	metricServerReqDur = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "api请求耗时",
		Help:    "API request metric",
		Buckets: []float64{5, 10, 25, 50, 100, 250, 500, 1000},
	},
		[]string{"path"},
	)

	metricServerReqCodeTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "api请求返回code统计",
		Help: "API request return code statics",
	},
		[]string{"path", "code"},
	)
)

// func RequestMetrics() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// startTime := time.Now()
// 		//
// 		// defer func() {
// 		// 	metricServerReqDur.
// 		// }()
// 	}
// }
