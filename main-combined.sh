#!/bin/sh

run-tests-junit-report.sh $1 > report.xml
ret=$?
cat report.xml
upload-test-artifacts.sh
exit $ret

