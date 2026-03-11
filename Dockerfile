FROM golang:alpine

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o crm.felix.com ./cmd/server

FROM scratch

COPY ./config /config

COPY --from=builder /build/crm.felix.com /

ENTRYPOINT ["/crm.felix.com", "config/local.yaml"]