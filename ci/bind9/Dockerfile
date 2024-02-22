FROM ubuntu/bind9:9.16-20.04_edge

RUN apt -y update \
    && apt -y install dnsutils iputils-ping net-tools

COPY db.com.gecko.local named.conf.local named.conf.options /etc/bind/
