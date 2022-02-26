#!/bin/bash
# Code generated .* DO NOT EDIT.
export REPO_ROOT_DIR:="$(git rev-parse --show-toplevel)/../board_generator"
set -e
export NAME="simulation"
export SYSTEM_NAMESPACE="simulation"
export CONFIG_PATH="${REPO_ROOT_DIR}/config/config.yaml"

cd `dirname $0`
docker network create simmulation-network || true
docker start sim-prometheus || docker run \
--name sim-prometheus \
--network simmulation-network \
-p 9090:9090 \
-v ${REPO_ROOT_DIR}/infra/config/prometheus.yaml:/etc/prometheus/prometheus.yml \
prom/prometheus:v2.23.0 &
docker start sim-jaeger || docker run \
--name sim-jaeger \
--network simmulation-network \
-e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
-p 5775:5775/udp \
-p 6831:6831/udp \
-p 6832:6832/udp \
-p 5778:5778 \
-p 16686:16686 \
-p 14268:14268 \
-p 14250:14250 \
-p 9411:9411 \
jaegertracing/all-in-one:1.21 &
../infra/scripts/server1.sh
docker stop sim-jaeger sim-prometheus
