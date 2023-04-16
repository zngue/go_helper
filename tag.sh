#!/bin/bash

version=${1:-"v1.0.26"}
git tag -d "${version}"
git push origin :refs/tags/"${version}"
msg=${2:-"Release ${version}"}
git tag -a "${version}" -m "${msg}"
git push origin "${version}"



