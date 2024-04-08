FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GGOPATH/src/github.com/SEANLEE0923/first-gin
COPY . $GGOPATH/src/github.com/SEANLEE0923/first-gin

RUN go build .

EXPOSE 2000
ENTRYPOINT [ "./first-gin" ]