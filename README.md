## Start minikube
Start docker desktop before
minikube start

## Create Deployment
kubectl create deployment hello-nginx --image=nginx:1.17-alpine

## Kubernets get
kubectl get pods
kubectl get deployments
kubectl get services

## kubernetes create service
kubectl expose deployment hello-nginx --type=LoadBalancer --port=80
minikube service hello-nginx

## Deploy Nginxx using yaml
kubectl delete deployments --all
kubectl delete deployment hello-nginx
kubectl apply -f configmap.yaml 
kubectl apply -f deployment.yaml 
minikube service nginx-service
kubectl exec hello-nginx-77cd7c6544-ggs7w -- apk add  bash
kubectl exec -ti hello-nginx-77cd7c6544-ggs7w -- bash

# Deploy MySQL k8s
kubectl create secret generic mysql-pass --from-literal=password='a1s2d3fa'
kubectl get secrets
kubectl apply -f mysql/deployment.yaml 
kubectl apply -f mysql/persistent-volume.yaml
kubectl get persistentvolumeclaim
kubectl get pods
kubectl exec -it mysql-server-7885f79678-xvlz7 bash
mysql -uroot -p

