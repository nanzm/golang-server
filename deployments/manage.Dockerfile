FROM golang:1.14.15-alpine3.13 AS builder
WORKDIR /app

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache tzdata
RUN apk add --no-cache git
ENV TZ Asia/Shanghai
ADD https://dl.nancode.cn/zoneinfo.zip /zoneinfo.zip

COPY ../go.mod .
COPY ../go.sum .
RUN GOPROXY="https://goproxy.cn,https://goproxy.io,direct" go mod download

COPY .. .
RUN ls -la
RUN env

RUN ["chmod", "+x", "/app/build.sh"]
RUN ["sh", "/app/build.sh"]
RUN cat config/version.go
RUN go build -o dora main.go


FROM alpine:3.10 AS final
WORKDIR /app
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/dora /app/
COPY --from=builder /zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO /zoneinfo.zip
ENTRYPOINT ["/app/dora"]
CMD ["manage"]