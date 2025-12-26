package server

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"

	"is-public-api/application/configs"

	"github.com/valyala/fasthttp"
)

type httpServer struct {
	config       *configs.ConfigServer
	server       *fasthttp.Server
	numInstances int
}

func NewHttpServer(conf *configs.ConfigServer) *httpServer {
	return &httpServer{
		config: conf,
		server: &fasthttp.Server{Name: "go", MaxRequestBodySize: 1024 * 1024 * 1024},
	}
}

func (httpSrv *httpServer) ListenAndServe(handler fasthttp.RequestHandler) {

	addr := fmt.Sprintf(":%d", httpSrv.config.Server.Port)

	prefork := flag.Bool("prefork", false, "use prefork")
	child := flag.Bool("child", false, "is child proc")
	flag.Parse()

	var ln net.Listener
	if *prefork {
		ln = doPrefork(*child, addr)
	} else {
		ln = getListener(addr)
	}

	httpSrv.server.Handler = handler

	log.Printf("HTTP server is listening on: %s\n", addr)
	if err := httpSrv.server.Serve(ln); err != nil {
		log.Fatalf("Error when serving incoming connections: %s", err)
	}
}

func (httpSrv *httpServer) SetInstances(numInstances int) {
	httpSrv.numInstances = numInstances
}

func (httpSrv *httpServer) getInstances() int {

	if httpSrv.numInstances > numCPU() {

		httpSrv.numInstances = numCPU()

	} else if httpSrv.numInstances <= 0 {

		httpSrv.numInstances = 1
	}
	return httpSrv.numInstances
}

func getListener(listenAddr string) net.Listener {
	ln, err := net.Listen("tcp4", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	return ln
}

func doPrefork(child bool, toBind string) net.Listener {
	var listener net.Listener
	if !child {
		addr, err := net.ResolveTCPAddr("tcp", toBind)
		if err != nil {
			log.Fatal(err)
		}
		tcplistener, err := net.ListenTCP("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}
		fl, err := tcplistener.File()
		if err != nil {
			log.Fatal(err)
		}
		len := numCPU()
		if len > 1 {
			len -= 1
		}

		children := make([]*exec.Cmd, len)
		for i := range children {
			children[i] = exec.Command(os.Args[0], "-prefork", "-child")
			children[i].Stdout = os.Stdout
			children[i].Stderr = os.Stderr
			children[i].ExtraFiles = []*os.File{fl}
			if err := children[i].Start(); err != nil {
				log.Fatal(err)
			}
		}
		for _, ch := range children {
			if err := ch.Wait(); err != nil {
				log.Print(err)
			}
		}
		os.Exit(0)
	} else {
		runtime.GOMAXPROCS(1)

		var err error
		listener, err = net.FileListener(os.NewFile(3, ""))
		if err != nil {
			log.Fatal(err)
		}
	}
	return listener
}

func numCPU() int {
	n := runtime.NumCPU()
	if n == 0 {
		n = 8
	}
	return n
}
