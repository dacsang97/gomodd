FROM golang:1.21-alpine
RUN go install github.com/cortesi/modd/cmd/modd@latest
WORKDIR /go
CMD ["modd"]