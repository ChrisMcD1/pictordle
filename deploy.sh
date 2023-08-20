#!/bin/bash


aws cloudformation update-stack --stack-name Pictordle --template-body file://template.yaml --capabilities CAPABILITY_IAM

