#!/bin/bash
set -euo pipefail

# Kubernetes Cluster Status Check
# Displays pods, services, and deployments across all namespaces.
# Useful for quick health checks and troubleshooting.

NAMESPACE="${1:---all-namespaces}"
[ "$NAMESPACE" != "--all-namespaces" ] && NAMESPACE="-n $NAMESPACE"

echo "=== Kubernetes Cluster Status ==="
echo "Time: $(date)"
echo "Context: $(kubectl config current-context)"
echo

# Cluster info
echo "--- Nodes ---"
kubectl get nodes -o wide

# Pods with status
echo -e "\n--- Pods ($NAMESPACE) ---"
kubectl get pods $NAMESPACE -o wide --sort-by='.metadata.namespace'

# Services
echo -e "\n--- Services ($NAMESPACE) ---"
kubectl get svc $NAMESPACE

# Deployments with ready replicas
echo -e "\n--- Deployments ($NAMESPACE) ---"
kubectl get deployments $NAMESPACE

# Show any pods not in Running/Succeeded state
echo -e "\n--- Problem Pods (not Running/Succeeded) ---"
kubectl get pods $NAMESPACE --field-selector='status.phase!=Running,status.phase!=Succeeded' 2>/dev/null || echo "None found"

# Resource usage (requires metrics-server)
echo -e "\n--- Resource Usage (top pods) ---"
kubectl top pods $NAMESPACE --sort-by=memory 2>/dev/null || echo "Metrics server not available"

echo -e "\n✓ Status check complete"
