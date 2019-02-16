FROM golang:latest

ADD . /go/src/github.com/cyantarek/ghost-clone

RUN cd /go/src/github.com/cyantarek/ghost-clone && go get -v

RUN go install github.com/cyantarek/ghost-clone

RUN cp -r /go/src/github.com/cyantarek/ghost-clone/keys /keys

RUN cp -r /go/src/github.com/cyantarek/ghost-clone/config /go/bin/config

ENTRYPOINT /go/bin/ghost-clone

EXPOSE 9000