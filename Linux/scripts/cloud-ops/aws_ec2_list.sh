#!/bin/bash
# List all EC2 instances with Name, ID, State, Type, and IP
# Usage: ./aws_ec2_list.sh [region]

REGION="${1:-us-east-1}"

aws ec2 describe-instances \
  --region "$REGION" \
  --query 'Reservations[*].Instances[*].[Tags[?Key==`Name`].Value|[0],InstanceId,State.Name,InstanceType,PublicIpAddress,PrivateIpAddress]' \
  --output table
