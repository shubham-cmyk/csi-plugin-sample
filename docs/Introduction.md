# Container Storage Interface (CSI) Plugin

## Capabilities of CSI Plugin

CSI plugins are equipped with the capabilities to manage the entire lifecycle of volumes in a Kubernetes cluster. The core capabilities include:

- **Dynamic provisioning and decommissioning of volumes:** Allows volumes to be created on-demand without admin intervention.

- **Attachment/mounting and detachment/unmounting of volumes:** Manages the connection between a volume and a running pod.

CSI plugins are primarily divided into two types:

- **Node Plugin:** A gRPC server that runs on each node where the storage provider volumes are provisioned. 

- **Controller Plugin:** A gRPC server that can run anywhere in the Kubernetes cluster. This manages the lifecycle of volumes.

## Services Implemented by CSI Plugin

CSI plugins implement the following services:

1. **Controller Service:** This service manages the lifecycle of volumes. It includes creating, deleting, attaching, and detaching volumes.

2. **Identity Service:** This service provides information about the plugin, including the versions it supports and its name.

3. **Node Service:** This service manages local volume operations like mounting and unmounting volumes.
