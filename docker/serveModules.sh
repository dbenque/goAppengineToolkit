#!/bin/bash

cmd="goapp serve --host=0.0.0.0"

for m in $MODULES
do
  ## extract the module name
  mname=${m##*/}
  echo "Preparing module $mname"

  ## prepare folder hosting the module
  mv $GOPATH/src/$m ~/project/$mname

done

## prepare the command to launch all modules
cmd="$cmd ./project/*/*.yaml"

echo "Launching server:"
echo $cmd
$cmd
