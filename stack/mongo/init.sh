#!/bin/bash -ex

cd "$(dirname "$0")"

openssl rand -base64 756 > mongo.key
