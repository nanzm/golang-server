FROM golang:1.16 AS builder

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
RUN go build -o transit cmd/transit/main.go


FROM alpine AS final
RUN apk update --no-cache
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

WORKDIR /cmd
COPY --from=builder /source/transit /cmd

CMD ["/cmd/transit", "-f", "/cmd/config.yml"]
