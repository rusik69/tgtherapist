FROM golang:latest as builder
WORKDIR /app
COPY . /app
RUN make build

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin/bin/tgtherapist-deploy-linux-amd64 /app
CMD ["./tgtherapist-deploy-linux-amd64"]