#!/bin/bash

#
# without code coverage profiling you can simply do
#     $ go test ./...
# to test all packages, but if you want to run coverage tests you have to call go test 
#  in every package folder separately
#

IFS=$'\n'
for adir in $(find . -type d -maxdepth 1 -not -path '*/\.*' -not -name '.')
do
	(cd "$adir" && go test -coverprofile .coverprofile)
done
unset IFS