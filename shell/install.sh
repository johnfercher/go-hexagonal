#!/usr/bin/env bash

go install golang.org/x/tools/cmd/goimports@latest
sudo cp $GOPATH/bin/goimports /usr/local/bin/

go install mvdan.cc/gofumpt@latest
sudo cp $GOPATH/bin/gofumpt /usr/local/bin/

curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0
sudo cp $GOPATH/bin/golangci-lint /usr/local/bin/

go install github.com/vektra/mockery/v2@latest
sudo cp $GOPATH/bin/mockery /usr/local/bin/

go install github.com/fe3dback/go-arch-lint@latest
sudo cp $GOPATH/bin/go-arch-lint /usr/local/bin