#!/bin/bash

cmd="goapp serve --host=0.0.0.0"

for m in $MODULES
do
  ## extract the module name
  mname=${m##*/}
  echo "Preparing module $mname"

  ## prepare folder hosting the module
  mkdir ~/project/$mname

  ## moving all module root files to the module path
  ## because the module cannot be in gopath: https://cloud.google.com/appengine/docs/go/#Go_Organizing_Go_apps
  find $GOPATH/src/$m -maxdepth 1 -type f -name "[^.]*" -exec mv {} ~/project/$mname \;

done

## prepare the command to launch all modules
cmd="$cmd ./project/*/*.yaml"


echo "Launching server:"
echo $cmd
$cmd
