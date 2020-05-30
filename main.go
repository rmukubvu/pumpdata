package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/rmukubvu/pumpdata/handles"
	"log"
	"net/http"
)

var port = flag.Int("p", 8002, "port number")

func main() {
	flag.Parse()
	r := handles.InitRouter()
	//convert to port format
	sPort := fmt.Sprintf(":%d", *port)
	//show on stdout
	fmt.Printf("Connecting to port [%s]", sPort)
	log.Fatal(http.ListenAndServe(sPort, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)))
}
