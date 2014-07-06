#!/bin/bash

#
# without code coverage profiling you can simply do
#     $ go test ./...
# to test all packages, but if you want to run coverage tests you have to call go test 
#  in every package folder separately
#

is_test_failed=0
IFS=$'\n'
for adir in $(find . -type d -maxdepth 1 -not -path '*/\.*' -not -name '.')
do
	(cd "$adir" && go test -coverprofile .coverprofile)
	if [ $? -ne 0 ]; then
		is_test_failed=1
	fi
done
unset IFS

echo "================================="
if [ $is_test_failed -ne 0 ]; then
	echo " [!] At least one of the tests failed!"
	exit 1
fi

echo " [OK] Tests finished with success"