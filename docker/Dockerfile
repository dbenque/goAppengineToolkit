# base image for Go Appengine project: dbenque/goappengine
FROM google/golang

MAINTAINER dbenque

RUN apt-get update

# Install cmake Python 2.x
RUN apt-get -y install procps build-essential cmake python-dev unzip

# Env
ENV HOME /home
WORKDIR /home

RUN curl -G https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_linux_amd64-1.9.19.zip > go_appengine_sdk_linux_amd64.zip && unzip -qq go_appengine_sdk_linux_amd64.zip && rm -rf go_appengine_sdk_linux_amd64.zip

ENV PATH $PATH:/home/go_appengine
ENV GOROOT /home/go_appengine/goroot
ENV GOPATH /home/go_appengine/gopath

RUN ln -s /home/go_appengine/goroot/bin/goapp /home/go_appengine/go

#### dont use localhost but 0.0.0.0
RUN sed -i "s/'localhost'/'0.0.0.0'/g" go_appengine/google/appengine/tools/devappserver2/devappserver2.py
RUN sed -i "s/'--api_port', type=PortParser(), default=0,/'--api_port', type=PortParser(), default=9000,/g" go_appengine/google/appengine/tools/devappserver2/devappserver2.py

EXPOSE 8080 8000 9000

# The GoAppengine project must not be under GOPATH : https://cloud.google.com/appengine/docs/go/#Go_Organizing_Go_apps
# So we move all root files of the project out of the GOPATH
RUN mkdir ~/project

# this will later on start all the modules that a child image will define in MODULES environment variable
ADD ./serveModules.sh ./goGetModules.sh /home/

# Run the project
#
# The modules must be given in the MODULES env variable, example:
# MODULES="github.com/dbenque/goAppengineToolkit/moduleData github.com/dbenque/goAppengineToolkit/moduleDefault github.com/dbenque/goAppengineToolkit/moduleHello github.com/dbenque/goAppengineToolkit/exampleApp"
#
CMD /home/goGetModules.sh && /home/serveModules.sh
