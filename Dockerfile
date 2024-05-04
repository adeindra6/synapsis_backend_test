FROM golang:1.18

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o synapsis-test

ENTRYPOINT ["/app/synapsis-test"]