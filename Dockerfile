FROM golang:1.14-alpine as builder
COPY . /app
WORKDIR /app
RUN go build -o out/hello-kubernetes

FROM alpine:3.12
COPY public /app/public
COPY static /app/static
COPY --from=builder /app/out/hello-kubernetes /app/hello-kubernetes

WORKDIR /app
EXPOSE 8080
CMD ["./hello-kubernetes"]