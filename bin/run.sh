#!/bin/sh
../build/ta-registry-v1.0.0-linux-amd64 --cert-path="../../etc/certs" \
		--logger-level="TRACE" \
		--local-db-path="../../var/db"
