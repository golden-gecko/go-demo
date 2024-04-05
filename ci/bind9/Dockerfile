FROM ubuntu/bind9:9.18-23.10_edge

RUN apt update -y \
    && apt install -y dnsutils iputils-ping net-tools

COPY db.com.gecko.local named.conf.local named.conf.options /etc/bind/
