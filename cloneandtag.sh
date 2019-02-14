#!/bin/sh

umask 077
DIR=$(mktemp -d)
trap "rm -rf $DIR" EXIT

git clone --depth 1 $1 $DIR > /dev/null 2>&1 && ctags -x --tag-relative=yes --c-kinds=f --language-force=C++ -f- -R $DIR | awk '{print $1 " " $4 " " $3}' | grep -P '\d+$'
