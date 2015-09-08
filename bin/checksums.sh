#!/bin/bash

export PKG_DIR=${PKG_DIR:=out}

if [[ "$(which shasum)X" == "X" ]]; then
  echo "shasum unavailable"
  exit 1
fi

find $PKG_DIR/* -exec shasum '{}' \;