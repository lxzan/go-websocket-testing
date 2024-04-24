package main

import (
	"github.com/antlabs/greatws"
	"github.com/lxzan/go-websocket-testing/internal"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
)

var (
	serverName = "greatws"
	upgrader   *greatws.UpgradeServer
)

func init() {
	internal.SetNumCPU()
	serverName = serverName + "-" + strings.ToLower(string(internal.AlphabetNumeric.Generate(6)))
}

func main() {
	var h = new(Handler)
	h.m = greatws.NewMultiEventLoopMust(
		greatws.WithEventLoops(runtime.NumCPU()), // 控制io go程数
		greatws.WithBusinessGoNum(240, 240, 240), // 控制业务go程数, 默认启动100个, 最小100个，最大10000个
		greatws.WithMaxEventNum(1000),
		greatws.WithLogLevel(slog.LevelError)) // epoll, kqueue
	h.m.Start()
	opt := []greatws.ServerOption{
		greatws.WithServerCallback(h),
		greatws.WithServerMultiEventLoop(h.m),
	}

	upgrader = greatws.NewUpgrade(opt...)
	lns := h.startServers(":8007")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
	for _, ln := range lns {
		ln.Close()
	}
}

type Handler struct {
	m *greatws.MultiEventLoop
}

func (c *Handler) startServers(addrs ...string) []net.Listener {
	lns := make([]net.Listener, 0, len(addrs))
	for _, addr := range addrs {
		mux := &http.ServeMux{}
		mux.HandleFunc("/connect", c.onWebsocket)
		server := http.Server{
			Handler: mux,
		}
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("Listen failed: %v", err)
		}
		lns = append(lns, ln)
		go func() {
			log.Printf("server exit: %v", server.Serve(ln))
		}()
	}
	return lns
}

func (c *Handler) onWebsocket(w http.ResponseWriter, r *http.Request) {
	_, err := upgrader.Upgrade(w, r)
	if err != nil {
		log.Printf("upgrade failed: %v", err)
		return
	}
}

func (c *Handler) OnOpen(conn *greatws.Conn) {

}

func (c *Handler) OnMessage(conn *greatws.Conn, opcode greatws.Opcode, bytes []byte) {
	_ = conn.WriteMessage(opcode, bytes)
}

func (c *Handler) OnClose(conn *greatws.Conn, err error) {

}
