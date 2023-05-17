package main

import (
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/lxzan/go-websocket-testing/internal"
	"log"
	"net/http"
	"strings"
)

var serverName = "gobwas"

func init() {
	internal.SetNumCPU()
	serverName = serverName + "-" + strings.ToLower(string(internal.AlphabetNumeric.Generate(6)))
}

func main() {
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

				_, _ = socket.Write(p)
			}
		}()
	})

	if err := http.ListenAndServe(":8003", nil); err != nil {
		log.Panic(err.Error())
	}
}
