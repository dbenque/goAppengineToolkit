M
goAppengineToolkit
==================

It is a set of tools that should help in developing application/modules for Go Appengine.
Developing but also testing locally and deploying.

Struct and interfaces
---------------------------

 - datastoreEntity
 This interface will help to store/retrieve/delete your object in the datastore. It takes care of the memcached layer and store the object in a json.

To be use this you object must implement the following interface:

    type DatastoreEntity interface {
      GetKey() string
      GetKind() string
    }

Then the available methods are:

    func Delete(context *appengine.Context, entity DatastoreEntity) error
    func Retrieve(context *appengine.Context, entity DatastoreEntity) error
    func Store(context *appengine.Context, entity DatastoreEntity) error

Local AppEngine
----------------
To easy the local testing of your go appengine application, a base Docker image is proposed. It contains everything to run a multiple module application.

Everything is made to ease the work for module that can be downloaded by "go get". Let's take the example proposed by [google in its tutorial](https://cloud.google.com/appengine/docs/go/#Go_Organizing_Go_apps):

    - project-dir
       - module1-dir
          app.yaml
          src1-1.go
          src1-2.go
       - module2-dir
          app.yaml
          src2-1.go
          src2-2.go
       - app-Dir
          dispatch.yaml

 1. Build the base **dbenque/goappengine** image. To do so, **go in folder goAppengineToolkit/docker** and type:

  `docker build --no-cache=true --rm=true -t "dbenque/goappengine" .`

 2. Run the application image, specifying the modules you would like to use:

  ` docker run --rm `
  ` -p 127.0.0.1:8080:8080 -p 127.0.0.1:8000:8000 -p 127.0.0.1:9000:9000 `
  `-e "MODULES=github.com/dbenque/goAppengineToolkit/moduleData github.com/dbenque/goAppengineToolkit/moduleDefault github.com/dbenque/goAppengineToolkit/moduleHello github.com/dbenque/goAppengineToolkit/exampleApp"`
  `dbenque/goappengine`

  If you are not using a Dispatch.yaml file and have several modules you may have to export more ports in that command (8081, 8082 ... )

  ### Module sub-directory structure: none!

  It is important that a module is represented by a FLAT list of go files. Any subdirectory structure must be done outside the module directory, and should be available in GOPATH. This is because:
  - Module dependencies will be retrieved by **go get** and placed in GOPATH
  - The prefered way of doing *import* is with [fully-qualified path][2]
  - [Go SDK scan both GOPATH and project paths][1]

  Thanks to this simple rule, we avoid conflict inside Go SDK.

  ### Local Development  

  If the module is under development it is possible to mount its local dev folder and gopath in the image instead of asking the container to *go get* it.
``
docker run --rm
-p 127.0.0.1:8080:8080 -p 127.0.0.1:8000:8000 -p 127.0.0.1:9000:9000
-v "$GOPATH/src/github.com/dbenque/goAppengineToolkit/moduleData:/home/project/moduleData"
-v "$GOPATH:/localgopath:ro" -e "LOCALGOPATH=`/bin/ls $GOPATH/src`"
 dbenque/goappengine
``

This command will mount you complete GOPATH. To be more selective place only the selected folders in the variable LOCALGOPATH.

  [1]: https://cloud.google.com/appengine/docs/go/#Go_Organizing_Go_apps
  [2]: https://golang.org/doc/code.html#remote
