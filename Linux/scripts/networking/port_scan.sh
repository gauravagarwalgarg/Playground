#!/bin/bash
# Quick port scan using /dev/tcp (no nmap needed)
# Usage: ./port_scan.sh <host> [start_port] [end_port]

HOST="${1:?Usage: $0 <host> [start] [end]}"
START="${2:-1}"
END="${3:-1024}"

echo "Scanning $HOST ports $START-$END..."
for ((port=$START; port<=$END; port++)); do
    (echo >/dev/tcp/$HOST/$port) 2>/dev/null && echo "  OPEN: $port"
done
echo "Scan complete."
