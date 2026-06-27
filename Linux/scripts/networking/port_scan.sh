#!/bin/bash
set -euo pipefail

# Port Scanner
# Checks if common ports are open on a target host.
# Uses /dev/tcp (bash built-in) with fallback to nc (netcat).

HOST="${1:-localhost}"
TIMEOUT=2

# Common ports to check
declare -A PORTS=(
    [22]="SSH"
    [80]="HTTP"
    [443]="HTTPS"
    [3306]="MySQL"
    [5432]="PostgreSQL"
    [6379]="Redis"
    [8080]="HTTP-Alt"
    [27017]="MongoDB"
    [9092]="Kafka"
    [2379]="etcd"
)

echo "=== Port Scan: ${HOST} ==="
echo "Time: $(date)"
echo

check_port_tcp() {
    local host=$1 port=$2
    (echo >/dev/tcp/"$host"/"$port") 2>/dev/null
}

check_port_nc() {
    local host=$1 port=$2
    nc -z -w "$TIMEOUT" "$host" "$port" 2>/dev/null
}

# Determine which method to use
if (echo >/dev/tcp/localhost/1) 2>/dev/null; then
    CHECKER="tcp"
elif command -v nc &>/dev/null; then
    CHECKER="nc"
else
    echo "Error: Neither /dev/tcp nor nc available"
    exit 1
fi

open_count=0
closed_count=0

printf "%-8s %-15s %s\n" "PORT" "SERVICE" "STATUS"
printf "%-8s %-15s %s\n" "----" "-------" "------"

for port in $(echo "${!PORTS[@]}" | tr ' ' '\n' | sort -n); do
    service="${PORTS[$port]}"
    if [ "$CHECKER" = "tcp" ]; then
        check_port_tcp "$HOST" "$port" && status="OPEN" || status="CLOSED"
    else
        check_port_nc "$HOST" "$port" && status="OPEN" || status="CLOSED"
    fi

    if [ "$status" = "OPEN" ]; then
        printf "%-8s %-15s \033[32m%s\033[0m\n" "$port" "$service" "$status"
        ((open_count++)) || true
    else
        printf "%-8s %-15s \033[31m%s\033[0m\n" "$port" "$service" "$status"
        ((closed_count++)) || true
    fi
done

echo
echo "Summary: ${open_count} open, ${closed_count} closed"
