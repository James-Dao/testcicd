FROM golang
maintainer james.xiong@daocloud.io
WORKDIR /gopath/app
ENV GOPATH /gopath/app
RUN apt-get update
RUN apt-get install -y vim
ADD . /gopath/app/
RUN go get github.com/codegangsta/cli
RUN go get github.com/gin-gonic/gin
RUN go install testcicd
ENV TZ Asia/Shanghai
EXPOSE 8080
CMD ["/gopath/app/bin/testcicd"]