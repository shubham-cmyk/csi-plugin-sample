#!/bin/bash

# Delete all Kubernetes resources
kubectl delete -f new-controller-plugin.yaml --force --grace-period=0
kubectl delete -f storageClass.yaml --force --grace-period=0
kubectl delete -f secret.yaml --force --grace-period=0
kubectl delete -f clusterRoleBinding.yaml --force --grace-period=0
kubectl delete -f clusterRole.yaml --force --grace-period=0
kubectl delete -f serviceAccount.yaml --force --grace-period=0
