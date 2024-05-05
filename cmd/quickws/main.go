package main

import (
	"github.com/antlabs/quickws"
	"github.com/lxzan/go-websocket-testing/internal"
	"log"
	"net/http"
	"strings"
	"time"
)

var serverName = "quickws"

func init() {
	internal.SetNumCPU()
	serverName = serverName + "-" + strings.ToLower(string(internal.AlphabetNumeric.Generate(6)))
}

func main() {
	handler := new(Handler)
	http.HandleFunc("/connect", func(writer http.ResponseWriter, request *http.Request) {
		socket, err := quickws.Upgrade(writer, request,
			quickws.WithServerCallback(handler),
			quickws.WithServerReadTimeout(5*time.Second),
		)
		if err != nil {
			log.Println(err.Error())
			return
		}
		go func() {
			socket.ReadLoop()
		}()
	})

	if err := http.ListenAndServe(":8007", nil); err != nil {
		log.Panic(err.Error())
	}
}

type Handler struct {
}

func (c *Handler) OnOpen(conn *quickws.Conn) {

}

func (c *Handler) OnMessage(conn *quickws.Conn, opcode quickws.Opcode, bytes []byte) {
	_ = conn.WriteMessage(opcode, bytes)
}

func (c *Handler) OnClose(conn *quickws.Conn, err error) {

}
