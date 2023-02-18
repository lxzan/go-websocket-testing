package main

import (
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/lxzan/go-websocket-testing/internal"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strings"
	"time"
)

var serverName = "gobwas"

func init() {
	serverName = serverName + "-" + strings.ToLower(string(internal.AlphabetNumeric.Generate(6)))
}

func main() {
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/connect", func(writer http.ResponseWriter, request *http.Request) {
		socket, _, _, err := ws.UpgradeHTTP(request, writer)
		if err != nil {
			return
		}

		go func() {
			defer socket.Close()

			for {
				p, _, err := wsutil.ReadClientData(socket)
				if err != nil {
					return
				}

				var t0 = time.Now()
				_, _ = socket.Write(p)
				latency := float64(time.Since(t0).Nanoseconds() / 1000)
				internal.LatencyDistributionCollector.WithLabelValues(serverName).Observe(latency)
				internal.LatencyPercentileCollector.WithLabelValues(serverName).Observe(latency)
			}
		}()
	})

	if err := http.ListenAndServe(":8003", nil); err != nil {
		log.Panic(err.Error())
	}
}
