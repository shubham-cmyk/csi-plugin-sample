package main

import (
	. "csi-plugin/pkg/driver"
	"flag"
	"fmt"
	"os"
)

func main() {

	// Take the input from the user i.e. endpoint token region
	var (
		endpoint = flag.String("endpoint", "", "Endpoint our gRPC server would run at")
		token    = flag.String("token", "", "Token of the storage provider")
		region   = flag.String("region", "ams3", "region wher the volumes are going to be provisioned")
	)

	// Parse the flags
	flag.Parse()

	// Read the Endpoint from the env Variable
	if *endpoint == "" {
		envEndpoint := os.Getenv("ENDPOINT")
		if envEndpoint != "" {
			endpoint = &envEndpoint
		} else {
			fmt.Println("Error: Endpoint not provided as an argument or environment variable.")
			return
		}
	}

	// Read the token from the env Variable
	if *token == "" {
		envToken := os.Getenv("TOKEN")
		if envToken != "" {
			token = &envToken
		} else {
			fmt.Println("Error: Token not provided as an argument or environment variable.")
			return
		}
	}

	driver, err := NewDriver(InputParams{
		Name:     DefaultName,
		Endpoint: *endpoint,
		Region:   *region,
		Token:    *token,
	})
	if err != nil {
		fmt.Printf("Error %s, creating new instance of driver", err.Error())
		return
	}

	driver.Run()
}
