#!/bin/bash
protoc owners.proto  --go_out=plugins=grpc:.