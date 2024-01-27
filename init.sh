#!/bin/bash

source .github/env.sh

go get -v -d
go install -v golang.org/x/mobile/cmd/gomobile@v0.0.0-20240213143359-d1f7d3436075
go install -v golang.org/x/mobile/cmd/gobind@v0.0.0-20240213143359-d1f7d3436075
gomobile init
