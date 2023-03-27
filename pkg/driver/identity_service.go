package driver

import (
	"context"
	"csi-plugin/logger"
	"csi-plugin/pkg"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (d *Driver) GetPluginInfo(context.Context, *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	logger.Info("GetPluginInfo RPC is called")
	return &csi.GetPluginInfoResponse{
		Name: pkg.DefaultName,
	}, nil

}
func (d *Driver) GetPluginCapabilities(context.Context, *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	logger.Info("GetPluginCapabilities RPC is called")
	return &csi.GetPluginCapabilitiesResponse{
		Capabilities: []*csi.PluginCapability{
			{
				Type: &csi.PluginCapability_Service_{
					Service: &csi.PluginCapability_Service{
						Type: csi.PluginCapability_Service_CONTROLLER_SERVICE,
					},
				},
			},
		}}, nil
}
func (d *Driver) Probe(context.Context, *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	logger.Info("Probe RPC is called")
	return &csi.ProbeResponse{
		Ready: &wrapperspb.BoolValue{
			Value: d.Ready,
		},
	}, nil
}
