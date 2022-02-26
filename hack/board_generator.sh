#!/bin/bash
set -e

export REPO_ROOT_DIR="$(git rev-parse --show-toplevel)/../board_generator"
export GRPC_SERVER_ADDRESS="localhost:50052"
export GRPC_CLIENT_ADDRESSES=""
export LOGGING_CONFIG_PATH="${REPO_ROOT_DIR}/config/logger.yaml"
export CONFIG_PATH="${REPO_ROOT_DIR}/config/config.yaml"
export JAEGER_SAMPLER_TYPE="ratelimiting"
export JAEGER_SAMPLER_PARAM="1.0"
export NAME="simulation"
export SYSTEM_NAMESPACE="simulation"

trap 'docker stop sim-timescaledb' SIGINT
docker network create simmulation-network || true
go run "../cmd/board_generator/main.go"
./sim.sh
