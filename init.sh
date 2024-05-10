#!/bin/bash

go get -v -d
go install -v golang.org/x/mobile/cmd/gomobile@v0.0.0-20240506190922-a1a533f289d3
go install -v golang.org/x/mobile/cmd/gobind@v0.0.0-20240506190922-a1a533f289d3
gomobile init
