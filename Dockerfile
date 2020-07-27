FROM leo2n/tesseractcentos:0.1
MAINTAINER leo2n<ttxs.an@gmail.com>


ENV MYPATH /usr/local
ENV TZ Asia/Shanghai

WORKDIR $MYPATH/teletraan

COPY mysql/mysql_config.json ./mysql/
COPY teletraanBin .
COPY public/teletraan.ico ./public/
COPY template/* ./template/
COPY README.md .
ENTRYPOINT ["/usr/local/teletraan/teletraanBin"]
EXPOSE 4001