FROM golang:1.17

RUN mkdir /new-server

ADD . /new-server

WORKDIR /new-server

RUN make setup
RUN make build

EXPOSE 3000

CMD ["./main"]
