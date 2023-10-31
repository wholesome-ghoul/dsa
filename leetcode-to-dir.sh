#!/bin/bash

leetcode_problem_url=$(basename "$1")
main=$leetcode_problem_url/main.go
main_test=$leetcode_problem_url/main_test.go

package=$(echo "package $leetcode_problem_url" | sed 's/-/_/g')

mkdir -p $leetcode_problem_url

touch $main $main_test

echo $package >>$main
echo $package >>$main_test
