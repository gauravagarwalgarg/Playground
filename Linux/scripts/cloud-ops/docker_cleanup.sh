#!/bin/bash
# Docker cleanup: remove stopped containers, dangling images, unused volumes
# Usage: ./docker_cleanup.sh [--dry-run]

set -euo pipefail

DRY_RUN="${1:-}"

echo "=== Docker Disk Usage ==="
docker system df

echo ""
echo "=== Stopped Containers ==="
docker ps -a --filter "status=exited" --format "{{.ID}} {{.Names}} ({{.Status}})"

echo ""
echo "=== Dangling Images ==="
docker images -f "dangling=true" --format "{{.ID}} {{.Repository}}:{{.Tag}} ({{.Size}})"

if [ "$DRY_RUN" = "--dry-run" ]; then
    echo ""
    echo "[DRY RUN] No changes made."
    exit 0
fi

echo ""
read -p "Proceed with cleanup? (y/N) " confirm
if [ "$confirm" = "y" ] || [ "$confirm" = "Y" ]; then
    docker container prune -f
    docker image prune -f
    docker volume prune -f
    docker network prune -f
    echo "Cleanup complete."
    docker system df
fi
