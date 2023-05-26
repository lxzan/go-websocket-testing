package main

import (
	"bytes"
	"github.com/gorilla/websocket"
	"github.com/lxzan/go-websocket-testing/internal"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

var serverName = "gorilla"

func init() {
	internal.SetNumCPU()
	serverName = serverName + "-" + strings.ToLower(string(internal.AlphabetNumeric.Generate(6)))
}

func main() {
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
			defer socket.Close()

			buf := make([]byte, 4000)
			payload := bytes.NewBufferString("")
			for {
				op, reader, err := socket.NextReader()
				if err != nil {
					return
				}

				payload.Reset()
				_, _ = io.CopyBuffer(payload, reader, buf)
				_ = socket.WriteMessage(op, payload.Bytes())
			}
		}()
	})

	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Panic(err.Error())
	}
}
