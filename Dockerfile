FROM debian:stable

RUN mkdir /pkg
COPY config.json /pkg
COPY Makefile /pkg
COPY go.mod /pkg
COPY cmd/. /pkg/cmd

RUN apt update
RUN apt install -y golang
RUN apt install -y make

WORKDIR /pkg
RUN make

CMD ["./helldivers2-bot"]
