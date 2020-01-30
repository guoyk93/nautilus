FROM golang:1.13 AS builder
WORKDIR /nautilus
ADD . .
RUN go build -mod vendor -o /usr/local/bin/svc_id nautilus/cmd/svc_id
RUN go build -mod vendor -o /usr/local/bin/svc_id_test nautilus/cmd/svc_id_test

FROM debian:10
ENV TZ=Asia/Shanghai
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y tzdata ca-certificates && \
    rm -rf /var/lib/apt/lists/* && \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY --from=builder /usr/local/bin/svc_id /usr/local/bin/svc_id
COPY --from=builder /usr/local/bin/svc_id_test /usr/local/bin/svc_id_test
ADD assets /assets