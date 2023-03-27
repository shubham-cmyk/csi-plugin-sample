package driver_test

import (
	"testing"
	"time"

	. "csi-plugin/pkg/driver"

	"github.com/digitalocean/godo"
	"google.golang.org/grpc"
	"k8s.io/apimachinery/pkg/util/wait"
)

const (
	testName     = "test-driver"
	testEndpoint = "unix:///tmp/csi.sock"
	testRegion   = "blr1"
	testToken    = "my-digitalocean-api-token"
)

func TestNewDriver(t *testing.T) {
	params := InputParams{
		Name:     testName,
		Endpoint: testEndpoint,
		Region:   testRegion,
		Token:    testToken,
	}
	d, err := NewDriver(params)
	if err != nil {
		t.Fatalf("Failed to create driver: %v", err)
	}
	if d.Name != testName {
		t.Errorf("Expected name %q, got %q", testName, d.Name)
	}
	if d.Endpoint != testEndpoint {
		t.Errorf("Expected endpoint %q, got %q", testEndpoint, d.Endpoint)
	}
	if d.Region != testRegion {
		t.Errorf("Expected region %q, got %q", testRegion, d.Region)
	}

	// Check that the storage and storageAction fields are set correctly.
	if d.Storage == nil {
		t.Error("storage field is nil")
	}
	if d.StorageAction == nil {
		t.Error("storageAction field is nil")
	}

	// Type assesrtion to do type check
	// value, ok := interfaceVariable.(ConcreteType)
	if _, ok := d.Storage.(*godo.StorageServiceOp); !ok {
		t.Error("storage field is not a *godo.StorageServiceOp")
	}

	if _, ok := d.StorageAction.(*godo.StorageActionsServiceOp); !ok {
		t.Error("storageAction field is not a *godo.StorageActionsServiceOp")
	}
}

func TestRun(t *testing.T) {
	params := InputParams{
		Name:     testName,
		Endpoint: testEndpoint,
		Region:   testRegion,
		Token:    testToken,
	}
	driver, err := NewDriver(params)
	if err != nil {
		t.Fatalf("Failed to create driver: %v", err)
	}

	// Start the driver in a separate goroutine
	go func() {
		err := driver.Run()
		if err != nil {
			t.Errorf("Failed to run driver: %v", err)
		}
	}()

	// Wait for the server to start with timeout of 3 second of wait
	err = wait.Poll(100*time.Millisecond, 3*time.Second, func() (bool, error) {
		if driver.Ready {
			return true, nil
		}
		t.Error("Driver not ready yet, waiting...")
		return false, nil
	})
	if err != nil {
		t.Errorf("Driver not ready after waiting: %v", err)
	}

	// Test that the driver is listening on the endpoint
	// We just dial a connection and note that we don't have any TLS on server
	conn, err := grpc.Dial(driver.Endpoint, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to connect to driver: %v", err)
	}

	// Test that the gRPC server can be stopped gracefully.
	if err := conn.Close(); err != nil {
		t.Errorf("Failed to stop gRPC server gracefully: %v", err)
	}

}
