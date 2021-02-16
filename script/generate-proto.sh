#!/bin/bash


PROTO_PATH=$(find . -name "idl" -print -quit)
TARGET_PATH=$PROTO_PATH/../proto
rm -rf "$TARGET_PATH"wq && mkdir -p "$TARGET_PATH"
protoc -I "$PROTO_PATH" --go_out=plugins=grpc:"$TARGET_PATH" "$PROTO_PATH"/**/*.proto
protoc -I "$PROTO_PATH" --go_out=plugins=grpc:"$TARGET_PATH" "$PROTO_PATH"/*.proto