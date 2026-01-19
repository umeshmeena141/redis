#!/bin/sh
#
# This script is used to compile your program 

set -e # Exit on failure
cd redis
go build -o build-redis-go redis
