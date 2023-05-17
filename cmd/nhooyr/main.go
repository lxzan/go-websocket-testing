package main

import (
	"context"
	"github.com/lxzan/go-websocket-testing/internal"
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
			for {
				op, p, err := socket.Read(context.Background())
				if err != nil {
					return
				}
				_ = socket.Write(context.Background(), op, p)
			}
		}()
	})

	if err := http.ListenAndServe(":8002", nil); err != nil {
		log.Panic(err.Error())
	}
}
