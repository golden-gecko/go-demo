FROM debian:bookworm-20240110-slim

RUN apt -y update \
    && apt -y install snapd \
    && snap install --classic certbot

RUN certbot --nginx
