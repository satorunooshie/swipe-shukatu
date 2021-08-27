FROM golang:1.16

ENV GOCS=linux
ENV GOACH=amd64
ENV CGO_ENABLED=0
ENV GO111MODULE=on

WORKDIR /go/src/github.com/satorunooshie/swipe-shukatu

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN make deps

EXPOSE 8888

CMD ["air", "-c", ".air.toml"]
