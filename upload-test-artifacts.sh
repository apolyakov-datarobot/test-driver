#!/bin/sh

testrunId="test-$(date +%d-%b-%y | tr A-Z a-z)-$(date +%s%3N)"
echo "Creating S3 bucket for test run $testrunId"
aws s3api create-bucket --bucket $testrunId --create-bucket-configuration LocationConstraint=us-west-1

for i in *.out report.xml; do
    echo "Uploading $i"
    aws s3 cp "$i" s3://$testrunId/ 
done
