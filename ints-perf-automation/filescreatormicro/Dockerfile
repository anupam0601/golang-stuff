FROM golang

RUN mkdir -p /go/src/github.com/anupam0601/golang-stuff/ints-perf-automation

ADD . /go/src/github.com/anupam0601/golang-stuff/ints-perf-automation
WORKDIR /go/src/github.com/anupam0601/golang-stuff/ints-perf-automation/filescreatormicro

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/anupam0601/golang-stuff/ints-perf-automation/filescreatormicro/cmd  -watch github.com/anupam0601/golang-stuff/ints-perf-automation/filescreatormicro
