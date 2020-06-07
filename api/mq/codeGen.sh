#!/bin/bash
protoc email.proto  --go_out=plugins=grpc:.  --proto_path=$HOME/Desktop/SELF/kuka/email-sender/api