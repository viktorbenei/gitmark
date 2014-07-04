#!/bin/bash

IFS=$'\n'
for adir in $(find . -type d -maxdepth 1 -not -path '*/\.*' -not -name '.')
do
	(cd "$adir" && go test)
done
unset IFS