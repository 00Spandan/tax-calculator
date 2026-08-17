#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
export PATH="${PATH}:$(go env GOPATH)/bin:${ROOT_DIR}/frontend/node_modules/.bin"

if ! command -v protoc >/dev/null 2>&1; then
  echo "protoc is required. Install with your package manager (e.g. apt-get install protobuf-compiler)."
  exit 1
fi

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.9
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

if ! command -v protoc-gen-ts_proto >/dev/null 2>&1; then
  echo "protoc-gen-ts_proto is required. Run: cd frontend && npm install"
  exit 1
fi

mkdir -p "${ROOT_DIR}/gen/go" "${ROOT_DIR}/gen/typescript"

protoc \
  --proto_path="${ROOT_DIR}/proto" \
  --go_out="${ROOT_DIR}/gen/go" \
  --go_opt=paths=source_relative \
  --go-grpc_out="${ROOT_DIR}/gen/go" \
  --go-grpc_opt=paths=source_relative \
  --ts_proto_out="${ROOT_DIR}/gen/typescript" \
  --ts_proto_opt=outputServices=generic-definitions,outputClientImpl=false,esModuleInterop=true,useOptionals=messages,snakeToCamel=true \
  "${ROOT_DIR}/proto/tax/tax.proto"
