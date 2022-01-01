package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	//flag is used for adding command line arguments
	port := flag.String("p", "8000", "port") //opens port 8000
	dir := flag.String("d", ".", "dir")      // . means check in the same folder
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*dir)))

	log.Printf("Serving %s on http port: %s\n", *dir, *port)

	log.Fatal(http.ListenAndServe(":"+*port, nil))

}
