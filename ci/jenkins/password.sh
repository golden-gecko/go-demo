#!/bin/bash -ex

cd "$(dirname "$0")"

docker exec -it demo-ci-jenkins-1 sh -c "cat /var/jenkins_home/secrets/initialAdminPassword"
