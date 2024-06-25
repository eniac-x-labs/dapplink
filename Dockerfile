FROM golang:stretch AS build-env

WORKDIR /go/src/github.com/eniac-x-labs/dapplink

RUN apt-get update -y
RUN apt-get install git -y

COPY . .

RUN make build

FROM golang:stretch

RUN apt-get update -y
RUN apt-get install ca-certificates jq -y

WORKDIR /root

COPY --from=build-env /go/src/github.com/eniac-x-labs/dapplink/build/dapplinkd /usr/bin/dapplinkd

EXPOSE 26656 26657 1317 9090

CMD ["dapplinkd"]
