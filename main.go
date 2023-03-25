package main

import (
	logger "csi-plugin/logger"
	. "csi-plugin/pkg"
	. "csi-plugin/pkg/driver"
	"flag"
	"os"
)

func main() {

	// Take the input from the user i.e. endpoint token region
	var (
		endpoint = flag.String("endpoint", "", "Endpoint our gRPC server would run at")
		token    = flag.String("token", "", "Token of the storage provider")
		region   = flag.String("region", "BLR1", "region wher the volumes are going to be provisioned")
	)

	// Parse the flags
	flag.Parse()

	// Read the Endpoint from the env Variable
	if *endpoint == "" {
		envEndpoint := os.Getenv("ENDPOINT")
		if envEndpoint != "" {
			endpoint = &envEndpoint
		} else {
			logger.Error("Endpoint not provided as an argument or environment variable.")
			return
		}
	}

	// Read the token from the env Variable
	if *token == "" {
		envToken := os.Getenv("TOKEN")
		if envToken != "" {
			token = &envToken
		} else {
			logger.Error("Token not provided as an argument or environment variable.")
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
		logger.Error("Error %s, creating new instance of driver", err.Error())
		return
	}

	// Start the Driver
	logger.Info("Starting the Driver")
	driver.Run()
}
