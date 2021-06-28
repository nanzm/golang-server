FROM golang:1.16.5-alpine3.13 AS builder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV TZ Asia/Shanghai

WORKDIR /source
COPY go.mod .
COPY go.sum .
RUN GOPROXY="https://goproxy.io,direct" go mod download


COPY . .
RUN ["chmod", "+x", "/source/build.sh"]
RUN ["sh", "/source/build.sh"]
RUN cat config/version.go
RUN go build -o manage cmd/manage/main.go


FROM alpine AS final
RUN apk update --no-cache
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

WORKDIR /cmd
COPY --from=builder /source/manage /cmd

CMD ["/cmd/manage", "-f", "/cmd/config.yml"]