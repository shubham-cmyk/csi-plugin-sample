# Sample CSI Plugin

This is a sample Container Storage Interface (CSI) plugin implemented in Go. The plugin provides basic storage functionality for containers and can be used as a starting point for building custom storage solutions.

## Code Structure

The codebase is organized into the following directories and files:

- `Dockerfile`: Contains instructions to build the CSI plugin as a Docker image.
- `logger/`: A custom logger package used for logging errors and info messages.
  - `log.go`: The main logger implementation file.
- `manifest/`: Contains Kubernetes YAML configuration files for deploying the CSI plugin.
  - `service_account.yaml`: Defines the ServiceAccount resource for the CSI plugin.
  - `storage_class.yaml`: Defines the StorageClass resource to provision volumes using this CSI plugin.
  - `cluster_role.yaml`: Defines the ClusterRole resource with necessary permissions for the CSI plugin.
  - `cluster_role_binding.yaml`: Binds the ClusterRole to the ServiceAccount.
- `pkg/driver/`: Contains core implementation files of the CSI plugin.
  - `driver.go`: Main driver implementation file that initializes and starts the gRPC server.
  - `identity_service.go`: Implements identity-related RPCs in the CSI spec, such as GetPluginInfo, GetPluginCapabilities, and Probe.
  - `controller_service.go`: Implements controller-related RPCs in the CSI spec, such as CreateVolume, DeleteVolume, ControllerPublishVolume, etc.
  - `node_service.go`: Implements node-related RPCs in the CSI spec, such as NodeStageVolume, NodeUnstageVolume, NodePublishVolume, etc.

At root level:

- `main.go`: Entry point of this application that sets up and starts the driver service
- `go.mod` and `go.sum` : Go module files containing dependencies information.

## Building the CSI Plugin

To build the CSI plugin, run:

```sh
go build -o sample-csi-plugin main.go

This command will generate a binary named sample-csi-plugin.


Building the Docker Image

To build the Docker image for the CSI plugin, run:


docker build -t your_dockerhub_username/sample-csi-plugin:latest .

Replace your_dockerhub_username with your actual Docker Hub username.


Deploying the CSI Plugin

To deploy the CSI plugin on a Kubernetes cluster, apply the manifest files in the manifest directory:


kubectl apply -f manifest/

This will create all necessary resources (ServiceAccount, StorageClass, ClusterRole, and ClusterRoleBinding) and deploy your sample CSI plugin on your Kubernetes cluster.
