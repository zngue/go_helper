#!/bin/bash
msg=${1:-"update"}
git add .
git commit -m "fix: ${msg}"
git push
