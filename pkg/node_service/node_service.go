package nodeservice

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
)

type node struct{}

var Node_service node

func (n *node) NodeStageVolume(context.Context, *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	return nil, nil
}
func (n *node) NodeUnstageVolume(context.Context, *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	return nil, nil
}
func (n *node) NodePublishVolume(context.Context, *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	return nil, nil
}
func (n *node) NodeUnpublishVolume(context.Context, *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	return nil, nil
}
func (n *node) NodeGetVolumeStats(context.Context, *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	return nil, nil
}
func (n *node) NodeExpandVolume(context.Context, *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	return nil, nil
}
func (n *node) NodeGetCapabilities(context.Context, *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	return nil, nil
}
func (n *node) NodeGetInfo(context.Context, *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	return nil, nil
}
