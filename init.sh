#!/bin/bash

source .github/env.sh

go get -v -d
go install -v golang.org/x/mobile/cmd/gomobile@v0.0.0-20231127183840-76ac6878050a
go install -v golang.org/x/mobile/cmd/gobind@v0.0.0-20231127183840-76ac6878050a
gomobile init
