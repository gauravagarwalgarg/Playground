#!/bin/bash
# System health snapshot: CPU, memory, disk, top processes
# Usage: ./system_health.sh

echo "=== System Health Report ==="
echo "Date: $(date)"
echo "Hostname: $(hostname)"
echo "Uptime: $(uptime -p)"
echo ""

echo "=== CPU ==="
echo "Load Average: $(cat /proc/loadavg | awk '{print $1, $2, $3}')"
echo "Cores: $(nproc)"
echo ""

echo "=== Memory ==="
free -h | grep -E "Mem|Swap"
echo ""

echo "=== Disk ==="
df -h | grep -E "^/dev|Filesystem"
echo ""

echo "=== Top 5 CPU Processes ==="
ps aux --sort=-%cpu | head -6
echo ""

echo "=== Top 5 Memory Processes ==="
ps aux --sort=-%mem | head -6
