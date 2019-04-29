# Hello Kubernetes

Simple React application which displays some container/pod related environment variables. 
It is one of the simplest applications you can deploy to test or demonstrate stuff on a Kubernetes Cluster.

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

## How to deploy on Kubernetes
```
kubectl apply -f ./kubernetes/
```

## How to run application
1. yarn install
2. yarn start


## TODO
- [ ] Add ingress.yaml.
- [ ] Show all available environment variables implicitly.