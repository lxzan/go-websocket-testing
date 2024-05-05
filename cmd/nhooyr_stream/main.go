package main

import (
	"context"
	"github.com/lxzan/go-websocket-testing/internal"
	"io"
	"log"
	"net/http"
	"nhooyr.io/websocket"
	"strings"
)

var serverName = "nhooyr_stream"

func init() {
	internal.SetNumCPU()
	serverName = serverName + "-" + strings.ToLower(string(internal.AlphabetNumeric.Generate(6)))
}

func main() {
	var options = &websocket.AcceptOptions{
		CompressionMode:    websocket.CompressionNoContextTakeover,
		InsecureSkipVerify: true,
	}

	http.HandleFunc("/connect", func(writer http.ResponseWriter, request *http.Request) {
		socket, err := websocket.Accept(writer, request, options)
		if err != nil {
			return
		}

		go func() {
			defer socket.Close(websocket.StatusNormalClosure, "sky is falling")
			buf := make([]byte, 4*1024)
			for {
				if err := readAndWrite(socket, buf); err != nil {
					return
				}
			}
		}()
	})

	if err := http.ListenAndServe(":8004", nil); err != nil {
		log.Panic(err.Error())
	}
}

func readAndWrite(socket *websocket.Conn, buf []byte) error {
	op, r, err := socket.Reader(context.Background())
	if err != nil {
		return err
	}
	w, err := socket.Writer(context.Background(), op)
	if err != nil {
		return err
	}
	_, err = io.CopyBuffer(w, r, buf)
	_ = w.Close()
	return err
}
