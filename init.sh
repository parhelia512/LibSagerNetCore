#!/bin/bash

source .github/env.sh

go get -v -d
go install -v golang.org/x/mobile/cmd/gomobile@v0.0.0-20231108233038-35478a0c49da
go install -v golang.org/x/mobile/cmd/gobind@v0.0.0-20231108233038-35478a0c49da
gomobile init
