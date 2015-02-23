
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

Local testing
----------------
To easy the local testing of your go appengine application, a base Docker image is proposed. It contains everything to run a multiple module application.

The image proposed here is not complete, it is a base image. Another image should be built on to of that one to run your application locally. Everything is made to ease the work for module that can be downloaded by "go get". Let's take the example proposed by [google in its tutorial](https://cloud.google.com/appengine/docs/go/#Go_Organizing_Go_apps):

    - project-dir
       - module1-dir
          app.yaml
          src1-1.go
          src1-2.go
       - module2-dir
          app.yaml
          src2-1.go
          src2-2.go

If the the 2 modules can be "go get" then it will be really easy to generate the docker image that will allow the local testing:

 1. Build the base **dbenque/goappengine** image. To do so, **go in folder goAppengineToolkit/docker** and type:

  `docker build -t "dbenque/goappengine" .`

 2. In your application environment, create a Dockerfile that declares the 2 modules and represents your application:
>  #Do no use cache in order to get fresh github sources
>  #docker build --no-cache -t "mylocalserver" .
>  
>  FROM dbenque/goappengine
>  MAINTAINER dbenque
>
>  #Setting the MODULES for "go get", modules must be separeted by space
>  ENV MODULES="github.com/mycount/module1 github.com/mycount/module2"
>
>  #Fetch the modules
>  RUN /home/goGetModules.sh
>
>  #Note: the final command is already defined in base image.

 3. Build your application image. Note the --no-cache that will force to fetch the latest available code for the modules.
 `docker build --no-cache -t "mylocalserver" .`


 4. Run the application image:
  ` docker run -p 127.0.0.1:8080:8080 -p 127.0.0.1:8000:8000 -p 127.0.0.1:9000:9000 mylocalserver`
