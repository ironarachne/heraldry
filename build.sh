#!/bin/bash
go build -o=build/makearms cmd/makearms/main.go
go build -o=build/armsapi cmd/armsapi/main.go