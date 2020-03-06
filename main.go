package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	flagBindTo = flag.String("bind-to", "localhost:9005", "bind to ")
	flagPath   = flag.String("path", "/dev/shm", "path to serve")
)

func main() {
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir(*flagPath)))
	log.Printf("Starting server on %s", *flagBindTo)
	log.Fatal(http.ListenAndServe(*flagBindTo, nil))
}
