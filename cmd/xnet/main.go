package main

import (
    "github.com/lxzan/go-websocket-testing/internal"
    "golang.org/x/net/websocket"
    "log"
    "net/http"
    "strings"
)

var serverName = "xnet"

func init() {
    internal.SetNumCPU()
    serverName = serverName + "-" + strings.ToLower(string(internal.AlphabetNumeric.Generate(6)))
}

func main() {
    srv := websocket.Server{
        Config: websocket.Config{},
        Handshake: func(config *websocket.Config, request *http.Request) error {
            return nil
        },
        Handler: Handler,
    }
    http.Handle("/connect", srv)

    if err := http.ListenAndServe(":8005", nil); err != nil {
        log.Panic(err.Error())
    }
}

func Handler(conn *websocket.Conn) {
    defer conn.Close()

    var buf = make([]byte, 16*1024)
    for {
        if err := websocket.Message.Receive(conn, &buf); err != nil {
            return
        }
        if err := websocket.Message.Send(conn, buf); err != nil {
            return
        }
    }
}
