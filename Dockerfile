FROM golang

ADD . /opt/easy-web-mock
WORKDIR /opt/easy-web-mock

#cound be changed to go run main.go -file /some/file.json
RUN /bin/bash