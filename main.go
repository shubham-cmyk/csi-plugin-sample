package main

import (
	"flag"
)

func main() {

	// Take the input from the user i.e. endpoint token region
	var (
		endpoint = flag.String("endpoint", "defaultValue", "Endpoint our gRPC server would run at")
		token    = flag.String("token", "defaultValue", "token of the storage provider")
		region   = flag.String("region", "ams3", "region wher the volumes are going to be provisioned")
	)

	// Parse the flags
	flag.Parse()

}
