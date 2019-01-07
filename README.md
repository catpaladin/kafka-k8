# Minikube-kafka example

Run Kafka on Minikube with Go Producers/Consumers. This repo assumes that all requirements are already installed.

#### Requirements
* Minikube
* virtualbox
* Docker

Setup Minikube
```
./minikube-start.sh
```

Create namespaces, storageclass, and RBAC
```
kubectl apply -f ./minikube
```

Create zookeeper
```
kubectl apply -f ./zookeeper
```

Create kafka
```
kubectl apply -f ./kafka
```

Create producers/consumers
```
kubectl apply -f ./pubsub
```
