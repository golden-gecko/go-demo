#!/bin/bash -ex

cd "$(dirname "$0")"

docker exec -it demo-ci-jenkins-1 sh -c "ssh-keygen -b 2048 -t rsa -f /var/jenkins_home/.ssh/id_rsa -q -N ''"
docker exec -it demo-ci-jenkins-1 sh -c "ssh-keyscan -H github.com >> ~/.ssh/known_hosts"

docker exec -it demo-ci-jenkins-1 sh -c "mkdir -p /var/lib/jenkins_home/jobs/demo-stack"
docker cp jobs/demo-stack/config.xml demo-ci-jenkins-1:/var/lib/jenkins_home/jobs/demo-stack/
