FROM golang:1.10

WORKDIR /go/src/app
COPY . .

RUN apt-get update
RUN apt-get install -y ca-certificates
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go install -v ./...

EXPOSE 3100

CMD ["app"]

# FROM ubuntu

# RUN \
#     apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 7F0CEB10 && \
#     echo 'deb http://downloads-distro.mongodb.org/repo/ubuntu-upstart dist 10gen' | sudo tee /etc/apt/sources.list.d/mongodb.list && \
#     apt-get update && \
#     apt-get install -y mongodb-org

# VOLUME ["/data/db"]
# WORKDIR /data

# EXPOSE 27017

# CMD ["mongod"]