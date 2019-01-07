## build instructions

change directory to `build`
```
cd build
```

build and push producer
```
docker build -f producer.Dockerfile -t <dockerhubuser>/producer:<version>
docker push <dockerhubuser>/producer:<version>
```

build and push consumer
```
docker build -f consumer.Dockerfile -t <dockerhubuser>/consumer:<version>
docker push <dockerhubuser>/consumer:<version>
```