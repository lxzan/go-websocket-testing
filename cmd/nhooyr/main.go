package main

import (
	"context"
	"github.com/lxzan/go-websocket-testing/internal"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"nhooyr.io/websocket"
	"strings"
	"time"
)

var serverName = "nhooyr"

func init() {
	serverName = serverName + "-" + strings.ToLower(string(internal.AlphabetNumeric.Generate(6)))
}

func main() {
	http.Handle("/metrics", promhttp.Handler())

	var options = &websocket.AcceptOptions{
		CompressionMode: websocket.CompressionDisabled,
	}

	http.HandleFunc("/connect", func(writer http.ResponseWriter, request *http.Request) {
		socket, err := websocket.Accept(writer, request, options)
		if err != nil {
			return
		}

		go func() {
			for {
				op, p, err := socket.Read(context.Background())
				if err != nil {
					return
				}

				var t0 = time.Now()
				_ = socket.Write(context.Background(), op, p)
				latency := float64(time.Since(t0).Nanoseconds() / 1000)
				internal.LatencyDistributionCollector.WithLabelValues(serverName).Observe(latency)
				internal.LatencyPercentileCollector.WithLabelValues(serverName).Observe(latency)
			}
		}()
	})

	if err := http.ListenAndServe(":8002", nil); err != nil {
		log.Panic(err.Error())
	}
}
