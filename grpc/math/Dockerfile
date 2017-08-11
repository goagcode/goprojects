FROM golang

COPY . /go/src/github.com/miguellgt/math
WORKDIR /go/src/github.com/miguellgt/math/server

RUN go get -d -v && go install github.com/miguellgt/math/server

ENTRYPOINT /go/bin/server

EXPOSE 50051
