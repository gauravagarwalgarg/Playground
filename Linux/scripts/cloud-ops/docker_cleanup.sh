#!/bin/bash
set -euo pipefail

# Docker Cleanup Script
# Stops all running containers, removes dangling images/volumes,
# and performs a system prune to reclaim disk space.

echo "=== Docker Cleanup ==="
echo "Starting at $(date)"
echo

# Stop all running containers
echo "--- Stopping all running containers ---"
running=$(docker ps -q)
if [ -n "$running" ]; then
    docker stop $running
    echo "Stopped $(echo "$running" | wc -l) container(s)"
else
    echo "No running containers"
fi

# Remove stopped containers
echo -e "\n--- Removing stopped containers ---"
docker container prune -f

# Remove dangling images (untagged)
echo -e "\n--- Removing dangling images ---"
docker image prune -f

# Remove unused volumes
echo -e "\n--- Removing unused volumes ---"
docker volume prune -f

# Remove unused networks
echo -e "\n--- Removing unused networks ---"
docker network prune -f

# Full system prune (optional: add --volumes for volume cleanup)
echo -e "\n--- System prune (unused data) ---"
docker system prune -f

# Show disk usage summary
echo -e "\n--- Current Docker disk usage ---"
docker system df

echo -e "\n✓ Cleanup complete at $(date)"
