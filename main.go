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
	proxy.Logger = log.New(os.Stdout, "", log.LstdFlags)

	proxy.Logger.Println("starting on 8080")

	log.Fatal(http.ListenAndServe(":8080", proxy))
}
