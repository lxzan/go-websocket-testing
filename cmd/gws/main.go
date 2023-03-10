package main

import (
	"github.com/lxzan/go-websocket-testing/internal"
	"github.com/lxzan/gws"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strings"
	"time"
)

var serverName = "gws"

func init() {
	serverName = serverName + "-" + strings.ToLower(string(internal.AlphabetNumeric.Generate(6)))
}

func main() {
	http.Handle("/metrics", promhttp.Handler())

	handler := new(Handler)

	upgrader := gws.NewUpgrader(func(c *gws.Upgrader) {
		c.EventHandler = handler
	})

	http.HandleFunc("/connect", func(writer http.ResponseWriter, request *http.Request) {
		socket, err := upgrader.Accept(writer, request)
		if err != nil {
			log.Println(err.Error())
			return
		}
		go socket.Listen()
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Panic(err.Error())
	}
}

type Handler struct {
	gws.BuiltinEventHandler
}

func (c *Handler) OnPing(socket *gws.Conn, payload []byte) {
	socket.WritePong(payload)
}

func (c *Handler) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()

	var t0 = time.Now()
	socket.WriteMessage(message.Opcode, message.Data.Bytes())
	//socket.WriteAsync(message.Opcode, message.Data.Bytes())
	latency := float64(time.Since(t0).Nanoseconds() / 1000)
	internal.LatencyDistributionCollector.WithLabelValues(serverName).Observe(latency)
	internal.LatencyPercentileCollector.WithLabelValues(serverName).Observe(latency)
}
