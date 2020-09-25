FROM golang:1.15.2-alpine3.12 AS build

RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/MiCasa-IoT/MainAPI
COPY go.mod go.sum ./
COPY cmd cmd
COPY configs configs
COPY internal internal

RUN GO111MODULE=on go build -o /go/bin/server ./cmd/main.go

FROM alpine:3.12
WORKDIR /usr/bin
COPY --from=build /go/bin .
COPY --from=build /go/src/github.com/MiCasa-IoT/MainAPI/configs configs
CMD /usr/bin/server
