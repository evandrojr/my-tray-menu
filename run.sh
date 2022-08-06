#!/usr/bin/env bash

yell() { echo "$0: $*" >&2; }
die() { yell "$*"; exit 111; }
try() { "$@" || die "cannot $*"; }

killall -9 __debug_bin
killall -9 my-tray-menu 
try go build
try ./my-tray-menu