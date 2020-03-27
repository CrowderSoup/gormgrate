#!/bin/sh
# Run the go-vet tool.
go vet ./..           || exit 1
# Run the test-cases, with race-detection.
go test -race ./...   || exit 1
# Everything passed, exit cleanly.
exit 0
