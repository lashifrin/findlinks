FROM golang:1.8.3 as builder
WORKDIR /Users/lls0605/work/go/src/findlinks/main
RUN go get -d -v github.com/golang/example/hello
COPY main/helloworld.go  .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o helloworld .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /Users/lls0605/work/go/src/findlinks/main/helloworld .
CMD ["./helloworld"]