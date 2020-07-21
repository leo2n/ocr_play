FROM leo2n/tesseractcentos:0.1
MAINTAINER leo2n<ttxs.an@gmail.com>


ENV MYPATH /usr/local
CMD mkdir -p $MYPATH/teletraan
WORKDIR $MYPATH/teletraan
CMD mkdir -p imageStore
EXPOSE 4001

COPY teletraanBin .
COPY public/* ./public/
COPY template/* ./template/
COPY README.md .
CMD ./teletraanBin