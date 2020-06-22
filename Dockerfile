FROM golang:latest AS builder

COPY . /go/src/github.com/alexapps/cloud-native-go
WORKDIR /go/src/github.com/alexapps/cloud-native-go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o cloud-native-go ./microservice.go

FROM scratch
ENV PORT 8083
EXPOSE 8083
COPY --from=builder /go/src/github.com/alexapps/cloud-native-go .
ENTRYPOINT ["/cloud-native-go"]
 