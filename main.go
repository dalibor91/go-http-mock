package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)
import "./requests"

func main() {

	var file = flag.String("file", "requests.json", "File with defined responses that mock server should return ")
	var host = flag.String("host", "0.0.0.0", "IP address or hostname for server to run")
	var port = flag.String("port", "8787", "Port which should be used to bind server")
	var debug= flag.String("debug", "__debug", "Endpoint used for debug")

	flag.Parse()

	if _, err := os.Stat(*file); os.IsNotExist(err) {
		log.Fatal(fmt.Sprintf("File '%s' not found", *file))
		os.Exit(1)
	}

	parsedResponses, err := requests.Load(*file)

	if err != nil {
		panic(err)
	}
	
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	signal.Notify(c, os.Kill, syscall.SIGKILL)
	go func() {
		<-c
		log.Println("Shutdown...")
		os.Exit(1)
	}()

	mockedResponses := parsedResponses.Build()

	http.HandleFunc("/", mockedResponses.GetHandler(*debug))

	log.Println(fmt.Sprintf("Listening for requests on %s:%s", *host, *port));

	http.ListenAndServe(fmt.Sprintf("%s:%s", *host, *port), nil)
}