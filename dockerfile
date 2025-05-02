FROM golang:1.24-alpine

RUN go version 

ENV GOPATH=/

COPY ./ ./ 

RUN go mod download
RUN go build -o trenlly-beta- ./cmd/main.go


CMD [ "./trenlly-beta-" ]