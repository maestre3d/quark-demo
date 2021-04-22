FROM golang:alpine as builder

LABEL "maintainer"="Alonso Ruiz <aruizea@neutrinocorp.org>"

RUN apk update && apk add --no-cache git

WORKDIR /go/src/github.com/maestre3d/quark-demo/user-service/

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o user ./cmd/http-server/main.go

FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/maestre3d/quark-demo/user-service .

EXPOSE 8081

CMD ["./user"]