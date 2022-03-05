#!/bin/bash

source .github/env.sh

go get -v -d
go install -v golang.org/x/mobile/cmd/gomobile@v0.0.0-20231006135142-2b44d11868fe
go install -v golang.org/x/mobile/cmd/gobind@v0.0.0-20231006135142-2b44d11868fe
gomobile init
