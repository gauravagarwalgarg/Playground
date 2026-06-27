#!/bin/bash
set -euo pipefail

# System Health Report
# Reports CPU, memory, disk usage with alerting thresholds.
# Shows top processes by CPU and memory consumption.

# --- Configuration ---
CPU_THRESHOLD=80     # Alert if CPU usage exceeds this %
MEM_THRESHOLD=85     # Alert if memory usage exceeds this %
DISK_THRESHOLD=90    # Alert if disk usage exceeds this %

echo "=== System Health Report ==="
echo "Host: $(hostname)"
echo "Time: $(date)"
echo "Uptime: $(uptime -p)"
echo

# --- CPU Usage ---
echo "--- CPU Usage ---"
cpu_idle=$(top -bn1 | grep "Cpu(s)" | awk '{print $8}' | cut -d. -f1)
cpu_usage=$((100 - cpu_idle))
echo "CPU Usage: ${cpu_usage}%"
if [ "$cpu_usage" -gt "$CPU_THRESHOLD" ]; then
    echo "⚠️  ALERT: CPU usage exceeds ${CPU_THRESHOLD}%!"
fi

# Load average
echo "Load Average: $(cat /proc/loadavg | awk '{print $1, $2, $3}')"
echo "CPU Cores: $(nproc)"

# --- Memory Usage ---
echo -e "\n--- Memory Usage ---"
mem_total=$(free -m | awk '/^Mem:/{print $2}')
mem_used=$(free -m | awk '/^Mem:/{print $3}')
mem_percent=$((mem_used * 100 / mem_total))
echo "Memory: ${mem_used}MB / ${mem_total}MB (${mem_percent}%)"
if [ "$mem_percent" -gt "$MEM_THRESHOLD" ]; then
    echo "⚠️  ALERT: Memory usage exceeds ${MEM_THRESHOLD}%!"
fi

# Swap
swap_total=$(free -m | awk '/^Swap:/{print $2}')
swap_used=$(free -m | awk '/^Swap:/{print $3}')
echo "Swap: ${swap_used}MB / ${swap_total}MB"

# --- Disk Usage ---
echo -e "\n--- Disk Usage ---"
df -h --type=ext4 --type=xfs --type=btrfs 2>/dev/null | head -20 || df -h | grep -E "^/dev/" | head -10
echo
# Check for any partition above threshold
while IFS= read -r line; do
    usage=$(echo "$line" | awk '{print $(NF-1)}' | tr -d '%')
    mount=$(echo "$line" | awk '{print $NF}')
    if [ "$usage" -gt "$DISK_THRESHOLD" ] 2>/dev/null; then
        echo "⚠️  ALERT: ${mount} at ${usage}% (exceeds ${DISK_THRESHOLD}%)"
    fi
done < <(df -h | grep -E "^/dev/")

# --- Top Processes ---
echo -e "\n--- Top 5 Processes by CPU ---"
ps aux --sort=-%cpu | head -6

echo -e "\n--- Top 5 Processes by Memory ---"
ps aux --sort=-%mem | head -6

echo -e "\n✓ Health check complete"
