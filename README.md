# Hello Kubernetes

Simple Go application which displays some container/pod related environment variables. 

## How to deploy on Kubernetes
```
kubectl apply -f ./kubernetes/
```

## How to run application
```
make run
```

## TODO
- [ ] Add ingress.yaml.
- [ ] Show all available environment variables implicitly.