# stage 0
FROM golang:latest as builder
WORKDIR /go/src/lander
COPY . .

RUN GOARCH=amd64 GOOS=linux go build -ldflags "-linkmode external -extldflags -static -w"

# stage 1
FROM scratch
WORKDIR /
COPY --from=builder /go/src/lander/lander .
COPY --from=builder /go/src/lander/template.html .
ENTRYPOINT ["/lander"]
