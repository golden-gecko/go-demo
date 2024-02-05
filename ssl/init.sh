#!/bin/bash -ex

cd "$(dirname "$0")"
cd certs

# CA (gecko)
openssl genrsa \
    -out ca-key.pem 2048

openssl req -new -x509 -nodes -days 365 \
    -subj "/C=PL/L=Cracow/O=GoldenGecko/CN=gecko" \
    -addext "subjectAltName = DNS:gecko" \
    -key ca-key.pem \
    -out ca-cert.pem

# Server (com.gecko)
openssl req -newkey rsa:2048 -nodes -days 365 \
    -subj "/C=PL/L=Cracow/O=GoldenGecko/CN=com.gecko" \
    -addext "subjectAltName = DNS:com.gecko" \
    -keyout server-key.pem \
    -out server-req.pem

openssl x509 -req -days 365 -set_serial 01 \
    -copy_extensions=copyall \
    -in server-req.pem \
    -out server-cert.pem \
    -CA ca-cert.pem \
    -CAkey ca-key.pem

# Client (*.com.gecko)
openssl req -newkey rsa:2048 -nodes -days 365 \
    -subj "/C=PL/L=Cracow/O=GoldenGecko/CN=*.com.gecko" \
    -addext "subjectAltName = DNS:*.com.gecko" \
    -keyout client-key.pem \
    -out client-req.pem

openssl x509 -req -days 365 -set_serial 01 \
    -copy_extensions=copyall \
    -in client-req.pem \
    -out client-cert.pem \
    -CA ca-cert.pem \
    -CAkey ca-key.pem

# Verify
openssl verify -CAfile ca-cert.pem \
    ca-cert.pem \
    server-cert.pem

openssl verify -CAfile ca-cert.pem \
    ca-cert.pem \
    client-cert.pem

# Chain
cat server-cert.pem ca-cert.pem > server-chain.pem
cat client-cert.pem server-cert.pem ca-cert.pem > client-chain.pem

# Copy
mkdir -p ../../ci/nginx/certs
mkdir -p ../../stack/nginx/certs
mkdir -p ../../stack/services/certs

cp *.pem ../../ci/nginx/certs
cp *.pem ../../stack/nginx/certs

cp ca-cert.pem ../../stack/services/certs
