FROM golang:1.11-stretch
ADD . /app/src/microservices-grpc-go-python
ENV GOPATH=/app
RUN go get -u google.golang.org/grpc \ && go get -u github.com/golang/protobuf/proto
WORKDIR /app/src/Proyecto2/servidor_Go
CMD ["go", "run", "main.go"]