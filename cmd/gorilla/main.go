package main

import (
	"github.com/gorilla/websocket"
	"github.com/lxzan/go-websocket-testing/internal"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

const serverName = "gorilla"

func main() {
	http.Handle("/metrics", promhttp.Handler())

	upgrader := websocket.Upgrader{
		EnableCompression: true,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	http.HandleFunc("/connect", func(writer http.ResponseWriter, request *http.Request) {
		socket, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			return
		}

		socket.SetPingHandler(func(appData string) error {
			return socket.WriteControl(websocket.PongMessage, nil, time.Now().Add(5*time.Second))
		})

		go func() {
			for {
				op, p, err := socket.ReadMessage()
				if err != nil {
					return
				}

				var t0 = time.Now()
				_ = socket.WriteMessage(op, p)
				latency := float64(time.Since(t0).Nanoseconds() / 1000)
				internal.LatencyDistributionCollector.WithLabelValues(serverName).Observe(latency)
				internal.LatencyPercentileCollector.WithLabelValues(serverName).Observe(latency)
			}
		}()
	})

	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Panic(err.Error())
	}
}
