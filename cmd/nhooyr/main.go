package main

import (
	"bytes"
	"context"
	"github.com/lxzan/go-websocket-testing/internal"
	"io"
	"log"
	"net/http"
	"nhooyr.io/websocket"
	"strings"
)

var serverName = "nhooyr"

func init() {
	internal.SetNumCPU()
	serverName = serverName + "-" + strings.ToLower(string(internal.AlphabetNumeric.Generate(6)))
}

func main() {
	var options = &websocket.AcceptOptions{
		CompressionMode: websocket.CompressionNoContextTakeover,
	}

	http.HandleFunc("/connect", func(writer http.ResponseWriter, request *http.Request) {
		socket, err := websocket.Accept(writer, request, options)
		if err != nil {
			return
		}

		go func() {
			buf := make([]byte, 4000)
			payload := bytes.NewBufferString("")

			for {
				op, reader, err := socket.Reader(context.Background())
				if err != nil {
					return
				}

				payload.Reset()
				_, _ = io.CopyBuffer(payload, reader, buf)
				_ = socket.Write(context.Background(), op, payload.Bytes())
			}
		}()
	})

	if err := http.ListenAndServe(":8002", nil); err != nil {
		log.Panic(err.Error())
	}
}
