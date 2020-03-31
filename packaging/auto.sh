#!/bin/bash

set -e

version=$(cat VERSION)
mkdir -p ~/rpmbuild/SOURCES

# update version
sed -i  -e "s/Version:.*/Version:\ ${version}/" ./packaging/*.spec

# link source
rm -rf ~/rpmbuild/SOURCES/openlan-ui-${version}
ln -s $(pwd) ~/rpmbuild/SOURCES/openlan-ui-${version}
