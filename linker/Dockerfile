FROM golang:1.22.3 AS builder

WORKDIR /app
ADD go.mod go.sum ./
RUN go mod download

ADD . .
RUN go build -o /app/linker

FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/linker /app/linker

EXPOSE 8001
CMD ["/app/linker"]