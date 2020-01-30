FROM golang:1.13 AS builder
WORKDIR /nautilus
ADD . .
RUN go build -mod vendor ./...

FROM debian:10
ENV TZ=Asia/Shanghai
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y tzdata ca-certificates && \
    rm -rf /var/lib/apt/lists/* && \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY --from=builder /nautilus/cmd/svc_id/svc_id /bin/svc_id
