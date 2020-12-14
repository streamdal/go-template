#!/bin/bash
set -e
if [ -f "/ca/batch_ca.crt" ]; then
  ln -s /ca/batch_ca.crt /usr/local/share/ca-certificates/batch_ca.crt
  update-ca-certificates
fi
exec "$@"
