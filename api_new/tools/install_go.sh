#!/bin/bash

set -ue -o pipefail

script_dir=$(dirname $0)
pj_root=${script_dir}/..

update_goenv () {
  goenv_path=${HOME}/.goenv

  if [ -e $goenv_path ]; then
    echo "update goenv"
    cd ${goenv_path}
    git pull
    cd -
  else
    echo "WARN: ${goenv_path} is not exists. please update goenv yourself"
  fi
}

install_go () {
  goversion_file=$pj_root/.go-version
  if [ ! -f ${goversion_file} ]; then
    echo "ERROR: ${goversion_file} is not exists"
    exit 1
  fi

  goversion=$(head $goversion_file)

  echo "install go ${goversion}. (Skip if the version appears to be installed already)"
  goenv install --skip-existing ${goversion}
}

update_goenv
install_go
