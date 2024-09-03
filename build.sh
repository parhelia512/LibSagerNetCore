#!/bin/bash

# quic-go needs asynctimerchan=1
CGO_LDFLAGS="-Wl,-z,max-page-size=16384" gomobile bind -v -androidapi 21 -trimpath -tags='disable_debug' -ldflags='-X=runtime.godebugDefault=asynctimerchan=1 -s -w -buildid=' . || exit 1
rm -r libcore-sources.jar

proj=../../app/libs
if [ -d $proj ]; then
  cp -f libcore.aar $proj
  echo ">> install $(realpath $proj)/libcore.aar"
fi
