FROM golang:1.22.3 AS builder

WORKDIR /app
ADD go.mod go.sum /app/

RUN go mod download

ADD . /app
RUN go build -o /app/server

FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/server /app/server

EXPOSE 8000
CMD ["/app/server"]
