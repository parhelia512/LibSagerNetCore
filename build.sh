#!/bin/bash

source .github/env.sh

gomobile bind -v -androidapi 21 -trimpath -tags='disable_debug' -ldflags='-s -w -buildid=' . || exit 1
rm -r libcore-sources.jar

proj=../SagerNet/app/libs
if [ -d $proj ]; then
  cp -f libcore.aar $proj
  echo ">> install $(realpath $proj)/libcore.aar"
fi
