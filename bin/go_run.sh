#!/bin/sh
go run ../main.go --cert-path="../../etc/certs" \
		--logger-level="TRACE" \
		--local-db-path="../../var/db"
