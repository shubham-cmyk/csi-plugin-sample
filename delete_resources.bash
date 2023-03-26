#!/bin/bash

# Delete all Kubernetes resources
kubectl delete -f manifest/controller-plugin.yaml --force --grace-period=0
kubectl delete -f manifest/node-plugin.yaml --force --grace-period=0
kubectl delete -f manifest/storageClass.yaml --force --grace-period=0
kubectl delete -f manifest/secret.yaml --force --grace-period=0
kubectl delete -f manifest/clusterRoleBinding.yaml --force --grace-period=0
kubectl delete -f manifest/clusterRole.yaml --force --grace-period=0
kubectl delete -f manifest/serviceAccount.yaml --force --grace-period=0

kubectl delete -f example/pod.yaml --force --grace-period=0
kubectl delete -f example/pvc.yaml --force --grace-period=0
