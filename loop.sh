#!/bin/bash
for i in {1..1000}
do
	echo "ROUND: #$i"
	go clean -testcache
	go test ./e2e -v --run TestPionE2ESimpleClientHelloHook
done


