FROM golang:1.14 AS builder
COPY . /app
WORKDIR /app
RUN go build -o monodiff cmd/monodiff/*.go

FROM alpine/git
COPY --from=builder /app/monodiff /usr/local/bin/monodiff
ENTRYPOINT ["monodiff"]
