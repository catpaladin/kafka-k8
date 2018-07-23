# kafka-k8

```
kubectl apply -f ./minikube
kubectl apply -f ./zookeeper
kubectl apply -f ./kafka
```

## Local Test

##### terminal 1
```
kubectl exec -it kafka-0 -n kafka

# unset port binding & run producer
unset JMX_PORT; bin/kafka-console-producer.sh --broker-list localhost:9092 --topic mytopic
```

##### terminal 2
```
ubectl exec -it kafka-0 -n kafka

# unset port binding & run consumer
unset JMX_PORT; bin/kafka-console-consumer.sh --zookeeper zookeeper.kafka.svc.cluster.local:2181 --topic mytopic
```
