#!/bin/bash

[ -d "$ANDROID_HOME" ] || ANDROID_HOME="$ANDROID_HOME"
[ -d "$ANDROID_HOME" ] || ANDROID_HOME="$HOME/Android/Sdk"
[ -d "$ANDROID_HOME" ] || ANDROID_HOME="$HOME/.local/lib/android/sdk"
[ -d "$ANDROID_HOME" ] || ANDROID_HOME="$HOME/Library/Android/sdk"

_NDK="$ANDROID_HOME/ndk/26.2.11394342"
[ -f "$_NDK/source.properties" ] || _NDK="$NDK"
[ -f "$_NDK/source.properties" ] || _NDK="$ANDROID_NDK_HOME"
[ -f "$_NDK/source.properties" ] || _NDK="$ANDROID_NDK_ROOT"
[ -f "$_NDK/source.properties" ] || _NDK="$ANDROID_NDK_LATEST_HOME"
[ -f "$_NDK/source.properties" ] || _NDK="$ANDROID_HOME/ndk-bundle"

if [ ! -f "$_NDK/source.properties" ]; then
  echo "Error: NDK not found."
  exit 1
fi

export ANDROID_HOME
export ANDROID_NDK_HOME=$_NDK
export NDK=$_NDK

if [ ! $(command -v go) ]; then
  if [ -d /usr/lib/go ]; then
    export PATH="$PATH:/usr/lib/go/bin"
  fi
fi

if [ $(command -v go) ]; then
  export PATH="$PATH:$(go env GOPATH)/bin"
fi