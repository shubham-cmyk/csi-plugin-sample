Here is a list of all the library and the framework required to build the project:

1. grpc[https://grpc.io/docs/] - You need to import this framework to build the grpc server and client. CSI plugin uses grpc to communicate with the kubernetes cluster.

2. library of your cloud provider that you would use to create/delete a Volume. For example, I am using digitalocean's library to create/delete a volume. You can find the library here:
https://pkg.go.dev/github.com/digitalocean/godo@v1.98.0
https://pkg.go.dev/github.com/digitalocean/go-metadata@v0.0.0-20220602160802-6f1b22e9ba8c

3. gRPC service are predefined in the proto file. You can find the proto file here: https://github.com/container-storage-interface/spec/blob/master/csi.proto


