#!/bin/bash

set -e

# Install creates the prometheus user and group using the
# name 'prometheus'. The prometheus user does not have a shell.
# This function can be called more than once as it is idempotent.
install() {
    username="prometheus"

    if getent group "$username" &>/dev/null; then
        echo "Group ${username} already exists."
    else
        groupadd "$username"
    fi

    if id "$username" &>/dev/null; then
        echo "User ${username} already exists"
        exit 0
    else
        useradd --shell /sbin/nologin --system "$username" -g "$username"
    fi
}

# Upgrade should perform the same steps as install
upgrade() {
    install
}

action="$1"

case "$action" in
  "0" | "install")
    install
    ;;
  "1" | "upgrade")
    upgrade
    ;;
  *)
    install
    ;;
esac
