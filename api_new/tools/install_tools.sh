#!/bin/bash
# this script is used to install all dependent tools
declare -a tools=(
  "github.com/MakotoNaruse/todocomment/cmd/todocomment@v1.0.0"
  "github.com/Songmu/make2help/cmd/make2help@v0.2.0"
  "github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.1"
  "github.com/pressly/goose/cmd/goose@v2.7.0"
  "golang.org/x/tools/cmd/godoc@v0.1.2"
  "golang.org/x/tools/cmd/goimports@v0.1.2"
  "github.com/golang/mock/mockgen@v1.6.0"
)

for package in "${tools[@]}"; do
  echo "installing" "$package"
  go install $package
done
