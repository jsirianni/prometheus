#!/bin/bash

set -e

rpm_install() {
    sudo yum remove -y prometheus
}

deb_install() {
    sudo apt-get remove -y prometheus
}

if command -v "dpkg" > /dev/null ; then
    deb_install
elif command -v "rpm" > /dev/null ; then
    rpm_install
else
    echo "failed to detect platform type"
    exit 1
fi
