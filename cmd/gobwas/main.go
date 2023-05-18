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
	http.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			return
		}
		go func() {
			defer conn.Close()

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					return
				}
				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					return
				}
			}
		}()
	})

	if err := http.ListenAndServe(":8003", nil); err != nil {
		log.Panic(err.Error())
	}
}
