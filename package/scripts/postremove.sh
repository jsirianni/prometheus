#!/bin/bash

set -e

# Remove deletes the prometheus systemd service file
# and reloads systemd. If the file does not exist, return early.
remove() {
    rm -f /usr/lib/systemd/system/prometheus.service || return
    systemctl daemon-reload
}

# Upgrade performs a no-op and is included here for future use.
upgrade() {
    return
}

action="$1"

case "$action" in
  "0" | "remove")
    remove
    ;;
  "1" | "upgrade")
    upgrade
    ;;
  *)
    remove
    ;;
esac
