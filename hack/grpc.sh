#!/bin/bash
REPO_ROOT_DIR="$(git rev-parse --show-toplevel)/../board_generator"
ROOT_DIR="~"

protoc --proto_path=${REPO_ROOT_DIR}/arrows/publisher/board/types \
  --go_out=${REPO_ROOT_DIR}/arrows/publisher/board/types \
  --go_opt=paths=source_relative \
  --go-grpc_out=${REPO_ROOT_DIR}/arrows/publisher/board/types \
  --go-grpc_opt=paths=source_relative \
  ${REPO_ROOT_DIR}/arrows/publisher/board/types/board.proto
