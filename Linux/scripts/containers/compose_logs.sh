#!/bin/bash
# Tail docker-compose logs with color and filtering
# Usage: ./compose_logs.sh [service] [lines]

SERVICE="${1:-}"
LINES="${2:-50}"

if [ -z "$SERVICE" ]; then
    docker compose logs --tail="$LINES" -f --timestamps
else
    docker compose logs "$SERVICE" --tail="$LINES" -f --timestamps
fi
