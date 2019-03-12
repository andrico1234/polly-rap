#!/bin/bash

BUCKET_NAME=$1
ERROR_DESCRIPTION="NoSuchBucket"
REL_PATH="../infra"

echo "Checking if bucket exists..."
if aws s3 ls "s3://$BUCKET_NAME" 2>&1 | grep -q $ERROR_DESCRIPTION; then
  echo "Creating bucket..."
  aws s3 mb s3://$BUCKET_NAME
else
  echo "Bucket exists!"
fi

echo "Packaging template"
aws cloudformation package \
  --template-file "$REL_PATH/master.yml" \
  --s3-bucket $BUCKET_NAME \
  --output-template-file "$REL_PATH/out/master-generate.yml"

echo "Deploying template"
aws cloudformation deploy \
  --template-file "$REL_PATH/out/master-generate.yml" \
  --s3-bucket $BUCKET_NAME \
  --capabilities CAPABILITY_IAM \
  --stack-name polly-stack
