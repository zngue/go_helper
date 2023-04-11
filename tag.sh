#!/bin/bash

version=${1:-"v1.0.20"}
msg=${2:-"Release ${version}"}
git tag -a "${version}" -m "${msg}"
git push origin "${version}"


