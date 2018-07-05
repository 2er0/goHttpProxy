package main

import (
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	"os"
)

func main() {
	httpProxy := goproxy.NewProxyHttpServer()
	httpProxy.Verbose = true
	httpProxy.Logger = log.New(os.Stdout, "", log.LstdFlags)

	log.Println("Starting proxy")

	log.Fatal(http.ListenAndServe(":8080", httpProxy))
}
