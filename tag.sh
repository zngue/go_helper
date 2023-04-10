#!/bin/bash
msg=${2:-"Release ${1}"}
version=${1:-"v1.0.20"}
git tag -a "${version}" -m "${msg}"
git push origin "${version}"

git add .

git commit -m "${msg}"

git push
