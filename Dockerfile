FROM golang:1.21

RUN mkdir /new-server

ADD . /new-server

WORKDIR /new-server

RUN make setup
RUN make build

EXPOSE 3000

HEALTHCHECK --interval=12s --timeout=12s --start-period=30s \
    CMD curl --fail http://localhost:3000/healthcheck || exit 1

CMD ["./main"]
