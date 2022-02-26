#!/bin/bash
REPO_ROOT_DIR="$(git rev-parse --show-toplevel)/../board_generator"
set -e
export NAME="build"
export SYSTEM_NAMESPACE="build"
export CONFIG_PATH="${REPO_ROOT_DIR}/config/config.yaml"
export JAEGER_SERVICE_NAME="server1"
# go run "${REPO_ROOT_DIR}/cmd/boards/main.go"
# ko resolve -f "${REPO_ROOT_DIR}/config/boards/kube_config" > "${REPO_ROOT_DIR}/config/boards/production_kube_config.yaml"
go run "${REPO_ROOT_DIR}/cmd/board_generator/main.go"
ko resolve -f "${REPO_ROOT_DIR}/config/kube/kube_config" > "${REPO_ROOT_DIR}/config/kube/production_kube_config.yaml"
echo "created at ${REPO_ROOT_DIR}/config/kube/production_kube_config.yaml"
