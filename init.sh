#!/bin/bash

go get -v -d
CGO_ENABLED=0 go install -v -trimpath -ldflags="-w -s" github.com/sagernet/gomobile/cmd/gobind@v0.1.4
CGO_ENABLED=0 go install -v -trimpath -ldflags="-w -s" github.com/sagernet/gomobile/cmd/gomobile@v0.1.4
gomobile init
