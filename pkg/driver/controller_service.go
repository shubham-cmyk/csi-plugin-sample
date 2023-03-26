package driver

import (
	"context"
	logger "csi-plugin/logger"
	"csi-plugin/pkg"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/digitalocean/godo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/apimachinery/pkg/util/wait"
)

func (d *Driver) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	logger.Info("CreateVolume RPC is called")

	// Check if Name Is Present in the request to create a volume
	if req.Name == "" {
		logger.Error("Can't create a volume without name got")
		return nil, status.Error(codes.InvalidArgument, "CreateVolume must be called with a req name")
	}

	// Get the required memory for the Volume Creation
	// make sure the value here is not less than or = 0
	// requriedBytes is not more than limiteBytes
	sizeBytes := req.CapacityRange.GetRequiredBytes()
	logger.Info("The Required volume size is %v", sizeBytes)
	// make sure volume capabilities have been specified
	if req.VolumeCapabilities == nil || len(req.VolumeCapabilities) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeCapabilities have not been specified")
	}
	// validate volume capabilities
	// make sure accessMode that has been specified by the PVC is actually supported by SP
	// make sure volumeMode that has been specified in the PVC is supported by us

	const gb = 1024 * 1024 * 1024
	// create the request struct
	volReq := godo.VolumeCreateRequest{
		Name:          req.Name,
		Region:        d.region,
		SizeGigaBytes: sizeBytes / gb,
	}

	// check if volumeContentSource is specified
	// if snapshot is specified, in that case, set the snapshot ID in the volume reqeust
	// you will also have to make sure that this snapshot is actually present
	// volReq.SnapshotID = req.VolumeContentSource.GetSnapshot().SnapshotId

	// if this user have not exceeded the limit
	// if this user can provision the requested amount etc

	// handle AccessibilityRequirements

	// actually call DO api to create the volume
	vol, res, err := d.storage.CreateVolume(ctx, &volReq)
	logger.Info("Got the response %v", res.StatusCode)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Failed provisoing the volume error %s\n", err.Error()))
	}

	return &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			CapacityBytes: sizeBytes,
			VolumeId:      vol.ID,
			// specify content source, but only in cases where its specified in the PVC
		},
	}, nil
}
func (d *Driver) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	logger.Info("DeleteVolume RPC is called")
	volumeID := req.VolumeId

	res, err := d.storage.DeleteVolume(ctx, volumeID)
	logger.Info("Got the response %v", res.StatusCode)
	if err != nil {
		logger.Error("Failed provisoing the volume error %s\n", err.Error())
		return nil, err
	}
	return &csi.DeleteVolumeResponse{}, nil
}
func (d *Driver) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	logger.Info("ControllerPublishVolume RPC is called")

	// check if volumeID is present and volume is available on SP
	if req.VolumeId == "" {
		return nil, status.Error(codes.InvalidArgument, "VolumeID is mandatory in ControllerPublishVolume request")
	}

	// check if NodeId is present and node is actually present on SP
	if req.NodeId == "" {
		return nil, status.Error(codes.InvalidArgument, "NodeID is mandatory in ControllerPublishVolume request")
	}

	vol, res, err := d.storage.GetVolume(ctx, req.VolumeId)
	logger.Info("Got the response %v", res.StatusCode)
	if err != nil {
		return nil, status.Error(codes.Internal, "Volume is not available anymore")
	}

	// check if the volume is already attached to a node
	// if it's attached to the correct node
	// or its attached to a wrong node

	// check readOnly is set, and you support readonly volumes
	// also check volumeCaps

	nodeID, err := strconv.Atoi(req.NodeId)
	if err != nil {
		logger.Error("was not able to convert nodeID to int value %s", err.Error())
		return nil, err
	}

	// Perform attach volume to the node
	action, res, err := d.storageAction.Attach(ctx, req.VolumeId, nodeID)
	logger.Info("Got the response %v", res.StatusCode)
	if err != nil {
		logger.Error("Failed to attach volume to the node %s", err.Error())
		return nil, err
	}

	// Wait For the attach volume Action to get Completed
	if err := d.waitForCompletionAttach(req.VolumeId, action.ID); err != nil {
		logger.Error("Waiting for volume to be attached got error: %s", err.Error())
		return nil, err
	}

	return &csi.ControllerPublishVolumeResponse{
		// Using Publish Context we are going to pass some key-value which may be need in later RPC
		PublishContext: map[string]string{
			pkg.VolNameKeyFromContPub: vol.Name,
		},
	}, nil
}
func (d *Driver) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	logger.Info("ControllerUnpublishVolume RPC is called")

	nodeID, err := strconv.Atoi(req.NodeId)
	if err != nil {
		logger.Error("was not able to convert nodeID to int value %s", err.Error())
		return nil, err
	}

	// Perform de-attach volume to the node
	action, res, err := d.storageAction.DetachByDropletID(ctx, req.VolumeId, nodeID)
	logger.Info("Got the response %v", res.StatusCode)
	if err != nil {
		logger.Error("Failed to de-attach volume to the node %s", err.Error())
		return nil, err
	}

	// Wait For the attach volume Action to get Completed
	if err := d.waitForCompletionDeAttach(req.VolumeId, action.ID); err != nil {
		logger.Error("Waiting for volume to be de-attached got error: %s", err.Error())
		return nil, err
	}

	return &csi.ControllerUnpublishVolumeResponse{}, nil
}
func (d *Driver) ValidateVolumeCapabilities(context.Context, *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	return nil, nil
}
func (d *Driver) ListVolumes(context.Context, *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	return nil, nil
}
func (d *Driver) GetCapacity(context.Context, *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	return nil, nil
}
func (d *Driver) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	logger.Info("ControllerGetCapabilities RPC is called")
	capability := []*csi.ControllerServiceCapability{}

	// The Current Capabililty we have allowed is create/delete and publish/unpublish of volume
	for _, c := range []csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
	} {
		capability = append(capability, &csi.ControllerServiceCapability{
			Type: &csi.ControllerServiceCapability_Rpc{
				Rpc: &csi.ControllerServiceCapability_RPC{
					Type: c,
				},
			},
		})
	}

	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: capability,
	}, nil
}
func (d *Driver) CreateSnapshot(context.Context, *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	return nil, nil
}
func (d *Driver) DeleteSnapshot(context.Context, *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	return nil, nil
}
func (d *Driver) ListSnapshots(context.Context, *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	return nil, nil
}
func (d *Driver) ControllerExpandVolume(context.Context, *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	return nil, nil
}
func (d *Driver) ControllerGetVolume(context.Context, *csi.ControllerGetVolumeRequest) (*csi.ControllerGetVolumeResponse, error) {
	return nil, nil
}

func (d *Driver) waitForCompletionAttach(volumeID string, actionID int) error {
	logger.Info("waitForCompletionAttach RPC is called")
	// Check in every 5 second whether the Volume is attached to node or not with timeout of 5 minute
	err := wait.Poll(5*time.Second, 5*time.Minute, func() (done bool, err error) {
		action, res, err := d.storageAction.Get(context.Background(), volumeID, actionID)
		logger.Info("Got the response %v", res.StatusCode)
		if err != nil {
			return false, nil
		}

		if action.Status == godo.ActionCompleted {
			return true, nil
		}
		return false, nil
	})
	return err
}

func (d *Driver) waitForCompletionDeAttach(volumeID string, actionID int) error {
	logger.Info("waitForCompletionDeAttach RPC is called")
	// Check in every 5 second whether the Volume is de-attached to node or not with timeout of 5 minute
	err := wait.Poll(5*time.Second, 5*time.Minute, func() (done bool, err error) {
		action, res, err := d.storageAction.Get(context.Background(), volumeID, actionID)
		// Http status not found of the volume confirm that volume is no longer present in provider
		logger.Info("Got the response %v", res.StatusCode)
		if err != nil && (res.StatusCode == http.StatusNotFound) {
			return true, nil
		}

		if action.Status == godo.ActionCompleted {
			return false, nil
		}
		return false, nil
	})
	return err
}
