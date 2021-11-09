FROM golang:1.15-alpine AS build
# Support CGO and SSL
RUN apk --no-cache add gcc g++ make
RUN apk add git
WORKDIR /go/src/application
COPY . .
ENV GOPATH="/go/src/application"
RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/handlers
RUN go get github.com/pusher/pusher-http-go
RUN GOOS=linux go build -ldflags="-s -w" -o main .

FROM alpine:3.10
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/application/main .
EXPOSE 80
ENTRYPOINT  ["./main"]