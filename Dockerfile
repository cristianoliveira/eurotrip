FROM golang:1.8

WORKDIR /go/
ADD . .

ADD ./service.sh src/

RUN ./service.sh dev_build

CMD ["./service.sh", "dev_run", "data/example"]

EXPOSE 8088
