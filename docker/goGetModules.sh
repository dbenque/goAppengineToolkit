#!/bin/bash

for m in $MODULES
do
  echo "Getting module $m"
  go get -d $m

done
