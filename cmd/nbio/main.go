package main

import (
	"compress/flate"
	"context"
	"flag"
	"fmt"
	"github.com/lesismal/nbio/nbhttp"
	"github.com/lesismal/nbio/nbhttp/websocket"
	"github.com/lxzan/go-websocket-testing/internal"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

var serverName = "nbio"

func init() {
	internal.SetNumCPU()
	serverName = serverName + "-" + strings.ToLower(string(internal.AlphabetNumeric.Generate(6)))
}

var (
	onDataFrame      = flag.Bool("UseOnDataFrame", false, "Server will use OnDataFrame api instead of OnMessage")
	errBeforeUpgrade = flag.Bool("error-before-upgrade", false, "return an error on upgrade with body")
)

func newUpgrader() *websocket.Upgrader {
	u := websocket.NewUpgrader()
	u.EnableCompression(true)
	u.SetCompressionLevel(flate.BestSpeed)
	u.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	u.OnOpen(func(conn *websocket.Conn) {
		conn.EnableWriteCompression(true)
	})

	if *onDataFrame {
		u.OnDataFrame(func(c *websocket.Conn, messageType websocket.MessageType, fin bool, data []byte) {
			// echo
			c.WriteFrame(messageType, true, fin, data)
		})
	} else {
		u.OnMessage(func(c *websocket.Conn, messageType websocket.MessageType, data []byte) {
			c.WriteMessage(messageType, data)
		})
	}

	u.OnClose(func(c *websocket.Conn, err error) {
	})
	return u
}

func onWebsocket(w http.ResponseWriter, r *http.Request) {
	if *errBeforeUpgrade {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("returning an error"))
		return
	}
	// time.Sleep(time.Second * 5)
	upgrader := newUpgrader()
	upgrader.BlockingModAsyncWrite = false
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	conn.SetReadDeadline(time.Time{})
	//fmt.Println("OnOpen:", wsConn.RemoteAddr().String())
}

func main() {
	flag.Parse()
	mux := &http.ServeMux{}

	mux.HandleFunc("/connect", onWebsocket)

	svr := nbhttp.NewServer(nbhttp.Config{
		IOMod:   nbhttp.IOModBlocking,
		Network: "tcp",
		Addrs:   []string{"0.0.0.0:8004"},
		Handler: mux,
	})

	err := svr.Start()
	if err != nil {
		fmt.Printf("nbio.Start failed: %v\n", err)
		return
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	svr.Shutdown(ctx)
}
