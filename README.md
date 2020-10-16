# Desafio Go Kubernetes

https://hub.docker.com/repository/docker/rvsiqueira/desafio-go-k8s
```
docker run rvsiqueira/desafio-go-k8s
```

## General commands used

1. Start minikube
```
Start docker desktop before
minikube start
```

2. Create Deployment
```
kubectl create deployment hello-nginx --image=nginx:1.17-alpine
```

3. Kubernets get
```
kubectl get pods
kubectl get deployments
kubectl get services
```

4. kubernetes create service
```
kubectl expose deployment hello-nginx --type=LoadBalancer --port=80
minikube service hello-nginx
```

5. Deploy Nginxx using yaml
```
kubectl delete deployments --all
kubectl delete deployment hello-nginx
kubectl apply -f configmap.yaml 
kubectl apply -f deployment.yaml 
minikube service nginx-service
kubectl exec hello-nginx-77cd7c6544-ggs7w -- apk add  bash
kubectl exec -ti hello-nginx-77cd7c6544-ggs7w -- bash
```

6. Deploy MySQL k8s
```
kubectl create secret generic mysql-pass --from-literal=password='a1s2d3fa'
kubectl get secrets
kubectl apply -f mysql/deployment.yaml 
kubectl apply -f mysql/persistent-volume.yaml
kubectl get persistentvolumeclaim
kubectl get pods
kubectl exec -it mysql-server-7885f79678-xvlz7 bash
mysql -uroot -p
```

7. Google Cloud Build local test
```
gcloud builds submit --config desafio-go/cloudbuild.yaml
```
