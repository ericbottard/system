#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

readonly root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." >/dev/null 2>&1 && pwd)

protoc -I ${root}/hack/ riff-serialization.proto --go_out=plugins=grpc:${root}/pkg/gateway/serialization
protoc -I ${root}/hack/ LiiklusService.proto --go_out=plugins=grpc:${root}/pkg/gateway/liiklus