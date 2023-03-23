package controllerservice

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
)

type controller struct{}

var Controller_service controller

func (c *controller) CreateVolume(context.Context, *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	return nil, nil
}
func (c *controller) DeleteVolume(context.Context, *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	return nil, nil
}
func (c *controller) ControllerPublishVolume(context.Context, *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	return nil, nil
}
func (c *controller) ControllerUnpublishVolume(context.Context, *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	return nil, nil
}
func (c *controller) ValidateVolumeCapabilities(context.Context, *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	return nil, nil
}
func (c *controller) ListVolumes(context.Context, *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	return nil, nil
}
func (c *controller) GetCapacity(context.Context, *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	return nil, nil
}
func (c *controller) ControllerGetCapabilities(context.Context, *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	return nil, nil
}
func (c *controller) CreateSnapshot(context.Context, *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	return nil, nil
}
func (c *controller) DeleteSnapshot(context.Context, *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	return nil, nil
}
func (c *controller) ListSnapshots(context.Context, *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	return nil, nil
}
func (c *controller) ControllerExpandVolume(context.Context, *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	return nil, nil
}
func (c *controller) ControllerGetVolume(context.Context, *csi.ControllerGetVolumeRequest) (*csi.ControllerGetVolumeResponse, error) {
	return nil, nil
}
