#!/usr/bin/env bash

protoc -I ./protocol protocol.proto --go_out=plugins=grpc:./protocol
