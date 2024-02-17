FROM golang:1.22-alpine as builder
WORKDIR /build
COPY . .
RUN go build -o homepage

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /build/homepage /app/homepage
USER nobody
ENTRYPOINT ["/app/homepage"]