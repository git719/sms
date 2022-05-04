#!/bin/bash
# install.sh
PrgPath="$(which $(basename $1))"
if [[ "$PrgPath" ]]; then
  printf "==> $1 is already installed at $PrgPath\n"
else
  cp -v $1 /usr/local/bin/
fi
