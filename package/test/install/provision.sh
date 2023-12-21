#!/bin/bash

set -e

rpm_install() {
    sudo bash /tmp/data/install-linux.sh --file /tmp/data/prometheus*linux*amd64.rpm
}

deb_install() {
    sudo bash /tmp/data/install-linux.sh --file /tmp/data/bindplane*linux*amd64.deb
}

start() {
    sudo systemctl start prometheus
}

if command -v "dpkg" > /dev/null ; then
    deb_install
elif command -v "rpm" > /dev/null ; then
    rpm_install
else
    echo "failed to detect platform type"
    exit 1
fi
start
