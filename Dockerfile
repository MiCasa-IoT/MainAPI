FROM golang:1.15.2-alpine3.12 AS build

RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/MiCasa-IoT/MainAPI
ENV SERVICE_NAME MiCasa-API
ENV APP /src/${SERVICE_NAME}/
ENV WORKDIR ${GOPATH}${APP}
COPY go.mod go.sum ./
COPY cmd cmd
COPY configs configs
COPY docs docs
COPY internal internal
COPY pkg pkg

RUN GO111MODULE=on go build -o /go/bin/server ./cmd/main.go

FROM alpine:3.12
WORKDIR /usr/bin
COPY --from=build /go/bin .
COPY --from=build /go/src/github.com/MiCasa-IoT/MainAPI/configs configs
CMD /usr/bin/server
