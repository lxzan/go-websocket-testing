package main

import (
	"github.com/lxzan/go-websocket-testing/internal"
	"github.com/lxzan/gws"
	"log"
	"net/http"
	"strings"
)

var serverName = "gws"

func init() {
	internal.SetNumCPU()
	serverName = serverName + "-" + strings.ToLower(string(internal.AlphabetNumeric.Generate(6)))
}

func main() {
	handler := new(Handler)

	upgrader := gws.NewUpgrader(handler, &gws.ServerOption{
		CompressEnabled: true,
	})

	http.HandleFunc("/connect", func(writer http.ResponseWriter, request *http.Request) {
		socket, err := upgrader.Upgrade(writer, request)
		if err != nil {
			log.Println(err.Error())
			return
		}
		go func() {
			socket.ReadLoop()
		}()
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
	socket.WriteMessage(message.Opcode, message.Bytes())
	_ = message.Close()
}
