FROM ubuntu/bind9:9.16-20.04_edge

RUN apt -y update \
    && apt -y install dnsutils iputils-ping net-tools

COPY db.com.gecko.local /etc/bind/db.com.gecko.local

COPY named.conf.local /etc/bind/named.conf.local
COPY named.conf.options /etc/bind/named.conf.options
