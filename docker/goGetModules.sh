#!/bin/bash

for m in $MODULES
do
  echo "Getting module $m"
  go get -d $m

done

for m in $LOCALGOPATH
do
  echo "Local GOPATH module $m"
  mkdir -p $GOPATH/src/$m
  rm -rf $GOPATH/src/$m
  ln -s /localgopath/src/$m $GOPATH/src/$m
done
