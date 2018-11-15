package main

import (
	"GoHello/routers"

	flag "github.com/spf13/pflag"
)

const (
	// default host and port
	defaultHost string = "localhost"
	defaultPort string = "8080"
)

func main() {
	// use router
	r := routers.Router()

	// server host
	host := defaultHost

	// server port
	port := defaultPort

	// use flag to get hostname and port parameter
	hostFlag := flag.StringP("hostname", "h", defaultHost, "The host to deploy at.")
	portFlag := flag.StringP("port", "p", defaultPort, "The port for listening.")
	flag.Parse()

	// if user enters the hostname and port he wants
	if len(*hostFlag) != 0 {
		host = *hostFlag
	}
	if len(*portFlag) != 0 {
		port = *portFlag
	}

	// run the router
	r.Run(host + ":" + port)
}
