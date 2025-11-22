#!/bin/bash -ex

cd "$(dirname "$0")"
cd ssl

openssl genrsa 2048 > ca-key.pem

openssl req -new -x509 -nodes -days 365000 -subj "/CN=com.gecko" \
   -key ca-key.pem \
   -out ca-cert.pem

openssl req -newkey rsa:2048 -nodes -days 365000 -subj "/CN=*.com.gecko" \
   -keyout server-key.pem \
   -out server-req.pem

openssl x509 -req -days 365000 -set_serial 01 \
   -in server-req.pem \
   -out server-cert.pem \
   -CA ca-cert.pem \
   -CAkey ca-key.pem

openssl req -newkey rsa:2048 -nodes -days 365000 -subj "/CN=registry.com.gecko" \
   -keyout client-key.pem \
   -out client-req.pem

openssl x509 -req -days 365000 -set_serial 01 \
   -in client-req.pem \
   -out client-cert.pem \
   -CA ca-cert.pem \
   -CAkey ca-key.pem

openssl verify -CAfile ca-cert.pem \
   ca-cert.pem \
   server-cert.pem

openssl verify -CAfile ca-cert.pem \
   ca-cert.pem \
   client-cert.pem

# Chain
cat client-cert.pem server-cert.pem ca-cert.pem > client-chain.pem
