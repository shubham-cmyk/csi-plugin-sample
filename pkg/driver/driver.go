package driver

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"path"
	"path/filepath"

	. "csi-plugin/pkg/controller_service"
	. "csi-plugin/pkg/identity_service"
	. "csi-plugin/pkg/node_service"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/digitalocean/godo"
	"google.golang.org/grpc"
)

type Driver struct {
	name     string
	region   string
	endpoint string

	srv *grpc.Server
	// http server, health check
	// storage clients
	storage       godo.StorageService
	storageAction godo.StorageActionsService

	ready bool
}

type InputParams struct {
	Name     string
	Endpoint string
	Token    string
	Region   string
}

const (
	DefaultName = "sample.csi.plugin"
)

func NewDriver(params InputParams) (*Driver, error) {
	if params.Token == "" {
		return nil, errors.New("token must be specified")
	}

	// client := godo.NewFromToken(params.Token)

	return &Driver{
		name:     params.Name,
		endpoint: params.Endpoint,
		region:   params.Region,
		// storage:       client.Storage,
		// storageAction: client.StorageActions,
	}, nil
}

// Start the gRPC server
func (d *Driver) Run() error {

	url, err := url.Parse(d.endpoint)
	if err != nil {
		log.Fatalf("Error parsing the endpoint: %s\n", err.Error())
		return err
	}

	if url.Scheme != "unix" {
		log.Fatalf("Only supported scheme is unix, but provided %s\n", url.Scheme)
		return fmt.Errorf("unsupported scheme")
	}

	grpcAddress := path.Join(url.Host, filepath.FromSlash(url.Path))
	if url.Host == "" {
		grpcAddress = filepath.FromSlash(url.Path)
	}

	if err := os.Remove(grpcAddress); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error removing listen address: %s\n", err.Error())
		return err
	}

	// Create a listener to Listen over the Unix Scheme
	listener, err := net.Listen("unix", grpcAddress)
	if err != nil {
		log.Fatalf("Got an Error while creating listener %v", err)
		return err
	}

	// Get a New GRPC server to Listen
	server := grpc.NewServer()

	// Register all services on the gRPC server
	csi.RegisterNodeServer(server, &Node_service)
	csi.RegisterControllerServer(server, &Controller_service)
	csi.RegisterIdentityServer(server, &Identity_service)

	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("Got an Error while starting server %v", err)
		return err
	}

	// Mark the Status of Driver to be ready if not found any error
	d.ready = true
	return nil
}
