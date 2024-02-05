#!/bin/bash -ex

cd "$(dirname "$0")"

curl --cacert ../../nginx/ssl/ca-cert.pem --ssl-no-revoke -sO https://registry.com.gecko/jnlpJars/agent.jar

java -jar agent.jar -jnlpUrl https://registry.com.gecko/computer/Windows/jenkins-agent.jnlp -secret d287188afb4178315691cd3a8dd957864c7c5b924f337471f7911d995c0822e9 -workDir "C:\Windows\Temp"
