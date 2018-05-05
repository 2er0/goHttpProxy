package main

import (
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	"os"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.Logger = log.New(os.Stdout, "Proxy", log.LstdFlags)

	log.Fatal(http.ListenAndServe(":8080", proxy))
}
