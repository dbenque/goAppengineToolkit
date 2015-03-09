#Do no use cache in order to get fresh github sources

#docker build --no-cache -t "mylocalserver" .

FROM dbenque/goappengine

MAINTAINER dbenque

#Setting the MODULES for "go get", modules must be separeted by space

ENV MODULES="github.com/dbenque/goAppengineToolkit/moduleData github.com/dbenque/goAppengineToolkit/moduleDefault github.com/dbenque/goAppengineToolkit/moduleHello github.com/dbenque/goAppengineToolkit/exampleApp"

#Fetch the modules

RUN /home/goGetModules.sh

#Note: the final command is already defined in base image.
