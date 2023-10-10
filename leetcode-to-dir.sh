#!/bin/bash

leetcode_problem_url=$(basename "$1")

mkdir -p $leetcode_problem_url

touch $leetcode_problem_url/main.go
touch $leetcode_problem_url/main_test.go
