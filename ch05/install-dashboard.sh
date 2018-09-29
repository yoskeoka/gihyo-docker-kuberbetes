#!/bin/bash

kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v1.8.3/src/deploy/recommended/kubernetes-dashboard.yaml
kubectl get pod --namespace=kube-system -l k8s-app=kubernetes-dashboard
kubectl proxy
open http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/
