#!/bin/bash -ex

cd "$(dirname "$0")"

mkdir -p /tmp/remoting

java -jar agent.jar -jnlpUrl https://jenkins.com.gecko/computer/WSL/jenkins-agent.jnlp -secret 19754c4a5e7390d6fec2b199f68099034e22f48d19f49aaca981eb90bbe5989f -workDir "/tmp" -failIfWorkDirIsMissing -noCertificateCheck
