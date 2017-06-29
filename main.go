package main

import (
	"log"
	"os"
)

func main() {
	os.Exit(Main())
}

func Main() int {
	term := make(chan os.Signal)

	webServer := NewWebServer()
	go webServer.Run()

	tailer := NewTailer()
	go tailer.Run()

	parser := NewParser(tailer.Out)
	go parser.Run()

	select {
	case <-term:
		log.Println("Received SIGTERM, exiting gracefully...")
	case err := <-webServer.ListenError():
		log.Println("Error starting web server, exiting gracefully:", err)
	case err := <-tailer.ListenError():
		log.Println("Error starting tail, exiting gracefully:", err)
	case err := <-parser.ListenError():
		log.Println("Error starting parser, exiting gracefully:", err)
	}
	return 0
}

//
// var requestOpts = []string{"method", "endpoint", "code"}
//
// var (
// 	httpRequestStatus2xx = prometheus.NewCounterVec(prometheus.CounterOpts{
// 		Name:        "http_requests_2xx_total",
// 		Help:        "Number status 2xx",
// 		ConstLabels: prometheus.Labels{},
// 	}, requestOpts)
// 	httpRequestStatus3xx = prometheus.NewCounterVec(prometheus.CounterOpts{
// 		Name:        "http_requests_3xx_total",
// 		Help:        "Number status 3xx",
// 		ConstLabels: prometheus.Labels{},
// 	}, requestOpts)
// 	httpRequestStatus4xx = prometheus.NewCounterVec(prometheus.CounterOpts{
// 		Name:        "http_requests_4xx_total",
// 		Help:        "Number status 4xx",
// 		ConstLabels: prometheus.Labels{},
// 	}, requestOpts)
// 	httpRequestStatus5xx = prometheus.NewCounterVec(prometheus.CounterOpts{
// 		Name:        "http_requests_5xx_total",
// 		Help:        "Number status 5xx",
// 		ConstLabels: prometheus.Labels{},
// 	}, requestOpts)
// 	httpRequestStatusxxx = prometheus.NewCounterVec(prometheus.CounterOpts{
// 		Name:        "http_requests_xxx_total",
// 		Help:        "Number status xxx",
// 		ConstLabels: prometheus.Labels{},
// 	}, requestOpts)
// 	httpResponseBodySize = prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 		Name:        "http_response_body_size_bytes",
// 		Help:        "Number status 5xx",
// 		ConstLabels: prometheus.Labels{},
// 	}, requestOpts)
// )
//
// func init() {
// 	prometheus.MustRegister(httpRequestStatus2xx)
// 	prometheus.MustRegister(httpRequestStatus3xx)
// 	prometheus.MustRegister(httpRequestStatus4xx)
// 	prometheus.MustRegister(httpRequestStatus5xx)
// 	prometheus.MustRegister(httpResponseBodySize)
// }
//
// func runWebServer(listenAddress, metricsPath string) {
// 	http.Handle("/metrics", promhttp.Handler())

// 	fmt.Printf("nginx log exporter is running on %s", listenAddress)
// 	panic(http.ListenAndServe(listenAddress, nil))
// }
//
// func main() {
// 	var (
// 		listenAddress = flag.String("web.listen-address", ":9393", "Address on which to expose metrics and web interface.")
// 		metricsPath   = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
// 		logPath       = flag.String("log.path", "/var/log/nginx/access.log", "Path under which to expose metrics.")
// 	)
// 	flag.Parse()
//
// 	go runWebServer(*listenAddress, *metricsPath)
//
// 	t, err := tail.TailFile(*logPath, tail.Config{Follow: true})
// 	if err != nil {
// 		panic(err)
// 	}
// 	parser := gonx.NewParser(`$remote_addr - $remote_user [$time_local] "$method $endpoint $http_version" $status $body_bytes_sent "$http_referer" "$http_user_agent"`)
// 	for line := range t.Lines {
// 		res, err := parser.ParseString(line.Text)
// 		if err != nil {
// 			panic(err)
// 		}
// 		status, err := res.Field("status")
// 		if err != nil {
// 			status = "xxx"
// 		}
// 		endpoint, _ := res.Field("endpoint")
// 		endpoint = strings.Split(endpoint, "?")[0]
//
// 		method, _ := res.Field("method")
//
// 		bodyBytesSent, _ := res.Field("body_bytes_sent")
//
// 		labels := []string{method, endpoint, status}
// 		bodyBytesSentF, _ := strconv.ParseFloat(bodyBytesSent, 64)
//
// 		httpResponseBodySize.WithLabelValues(labels...).Set(bodyBytesSentF)
// 		switch string(status[0]) {
// 		case "1":
//
// 		case "2":
// 			httpRequestStatus2xx.WithLabelValues(labels...).Inc()
// 		case "3":
// 			httpRequestStatus3xx.WithLabelValues(labels...).Inc()
// 		case "4":
// 			httpRequestStatus4xx.WithLabelValues(labels...).Inc()
// 		case "5":
// 			httpRequestStatus5xx.WithLabelValues(labels...).Inc()
// 		}
//
// 	}
//
// }
