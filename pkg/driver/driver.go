package driver

import (
	logger "csi-plugin/logger"
	"errors"
	"net"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/digitalocean/godo"
	"google.golang.org/grpc"
)

type Driver struct {
	name     string
	region   string
	endpoint string

	srv   *grpc.Server // grpc Server
	ready bool         //  health check

	storage       godo.StorageService // storage provider client
	storageAction godo.StorageActionsService
}

type InputParams struct {
	Name     string
	Endpoint string
	Token    string
	Region   string
}

func NewDriver(params InputParams) (*Driver, error) {
	if params.Token == "" {
		return nil, errors.New("token must be specified")
	}

	// Create a Client that could interace with the provider
	client := godo.NewFromToken(params.Token)

	return &Driver{
		name:          params.Name,
		endpoint:      params.Endpoint,
		region:        params.Region,
		storage:       client.Storage,
		storageAction: client.StorageActions,
	}, nil
}

// Start the gRPC server
func (d *Driver) Run() error {

	url, err := url.Parse(d.endpoint)
	if err != nil {
		logger.Error("Error parsing the endpoint: %s\n", err.Error())
		return err
	}

	if url.Scheme != "unix" {
		logger.Error("Only supported scheme is unix, but provided %s\n", url.Scheme)
		return errors.New("only supported scheme is unix")
	}

	grpcAddress := path.Join(url.Host, filepath.FromSlash(url.Path))
	if url.Host == "" {
		grpcAddress = filepath.FromSlash(url.Path)
	}

	if err := os.Remove(grpcAddress); err != nil && !os.IsNotExist(err) {
		logger.Error("Error removing listen address: %s\n", err.Error())
		return err
	}

	// Create a listener to Listen over the Unix Scheme
	listener, err := net.Listen("unix", grpcAddress)
	if err != nil {
		logger.Error("Got an Error while creating listener %v", err)
		return err
	}

	// Get a New GRPC server to Listen
	server := grpc.NewServer()

	// Register all services on the gRPC server
	csi.RegisterNodeServer(server, d)
	csi.RegisterControllerServer(server, d)
	csi.RegisterIdentityServer(server, d)

	// Starting the Server
	logger.Info("Starting the gRPC server")
	err = server.Serve(listener)
	if err != nil {
		logger.Error("Got an Error while starting gRPC server %v", err)
		return err
	}

	// Mark the Status of Driver to be ready if not found any error
	d.ready = true
	return nil
}
