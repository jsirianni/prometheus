#!/bin/bash

# Install handles systemd service management. This function
# can be called more than once as it is idempotent.
install() {
    systemctl daemon-reload
}

# Upgrade performs the same steps as install.
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
