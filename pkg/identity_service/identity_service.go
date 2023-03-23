package identityservice

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
)

type identity struct{}

var Identity_service identity

func (i *identity) GetPluginInfo(context.Context, *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	return nil, nil
}
func (i *identity) GetPluginCapabilities(context.Context, *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	return nil, nil
}
func (i *identity) Probe(context.Context, *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	return nil, nil
}
