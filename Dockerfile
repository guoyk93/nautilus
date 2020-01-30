FROM golang:1.13 AS builder
WORKDIR /nautilus
ADD . .
RUN go build -mod vendor -o /usr/local/bin/svc_id nautilus/cmd/svc_id
RUN go build -mod vendor -o /usr/local/bin/svc_id_test nautilus/cmd/svc_id_test
RUN go build -mod vendor -o /usr/local/bin/web_main nautilus/cmd/web_main
RUN go build -mod vendor -o /usr/local/bin/web_null nautilus/cmd/web_null

FROM debian:10
ENV TZ=Asia/Shanghai
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y tzdata ca-certificates && \
    rm -rf /var/lib/apt/lists/* && \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY --from=builder /usr/local/bin/svc_id /usr/local/bin/svc_id
COPY --from=builder /usr/local/bin/svc_id_test /usr/local/bin/svc_id_test
COPY --from=builder /usr/local/bin/web_main /usr/local/bin/web_main
COPY --from=builder /usr/local/bin/web_null /usr/local/bin/web_null
ADD assets /assets