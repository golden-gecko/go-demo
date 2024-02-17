#!/bin/bash -ex

cd "$(dirname "$0")"

nomad job run -detach jobs/api.hcl
nomad job run -detach jobs/mongo.hcl
nomad job run -detach jobs/node.hcl
nomad job run -detach jobs/roach.hcl
