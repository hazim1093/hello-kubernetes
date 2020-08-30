build: 
	go build -o out/hello-kubernetes

run: build
	./out/hello-kubernetes