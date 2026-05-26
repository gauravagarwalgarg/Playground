#!/bin/bash
# Quick Kubernetes pod health check
# Usage: ./k8s_pod_status.sh [namespace]

NAMESPACE="${1:---all-namespaces}"

if [ "$NAMESPACE" != "--all-namespaces" ]; then
    NAMESPACE="-n $NAMESPACE"
fi

echo "=== Pods Not Running ==="
kubectl get pods $NAMESPACE --field-selector=status.phase!=Running 2>/dev/null || echo "None"

echo ""
echo "=== Pod Restarts (>3) ==="
kubectl get pods $NAMESPACE -o json 2>/dev/null | \
  jq -r '.items[] | select(.status.containerStatuses[]?.restartCount > 3) | "\(.metadata.namespace)/\(.metadata.name) restarts=\(.status.containerStatuses[0].restartCount)"' 2>/dev/null || echo "None"

echo ""
echo "=== Node Status ==="
kubectl get nodes -o wide 2>/dev/null || echo "kubectl not configured"
