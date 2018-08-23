FROM golang:1.10.3-alpine3.8
MAINTAINER ali
ENV SOURCES /go/src/SimpleAPIBasePlatform/Doc/Docker/SimpleAPi/
COPY . ${SOURCES}
RUN cd ${SOURCES} && go install
ENTRYPOINT SimpleAPi