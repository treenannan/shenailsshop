FROM golang:1.14.4-alpine3.11
RUN apk update && apk add gcc git
ENV GO111MODULE=on
WORKDIR /
COPY /src/go.mod .
COPY /src/go.sum .
RUN go mod download
COPY . .
WORKDIR /src/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ./shenails.com

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 / .
CMD ["./shenails.com"]