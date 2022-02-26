#!/bin/bash
# Code generated .* DO NOT EDIT.

cd `dirname $0`
export NAME="board_generator"
export SYSTEM_NAMESPACE="board_generator"
export JAEGER_SERVICE_NAME="board_generator"
export PROMETHEUS_ADDRESS="${PROMETHEUS_ADDRESS:-localhost:9000}"

go run ../cmd/board_generator/main.go
