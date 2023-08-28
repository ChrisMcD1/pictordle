#!/bin/bash


# Build and push app server to ECR
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 101357155028.dkr.ecr.us-east-1.amazonaws.com
docker build ./app-server -t 101357155028.dkr.ecr.us-east-1.amazonaws.com/pictordle:latest
docker push 101357155028.dkr.ecr.us-east-1.amazonaws.com/pictordle:latest

# Deploy cloud formation stack. (Needs ECR to already be updated)
aws cloudformation deploy --stack-name Pictordle --template-file ./template.yaml --capabilities CAPABILITY_NAMED_IAM

