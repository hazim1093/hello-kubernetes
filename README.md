# Hello Kubernetes

Simple Go application which displays some container/pod related environment variables. 

## How to deploy on Kubernetes
```
kubectl apply -f ./kubernetes/
```

## How to run application
```
go run main.go
```

## TODO
- [ ] Add ingress.yaml.
- [ ] Show all available environment variables implicitly.