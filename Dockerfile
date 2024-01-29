################# PROD ######################

FROM golang:1.20.13-alpine AS builder
ARG VERSION=PROD
COPY . /server
WORKDIR /server
ENV GO111MODULE=on
RUN CGO_ENABLED=0 go build -ldflags "-X 'github.com/asciiflix/server/config.Version=$VERSION'" -o /main .

FROM alpine:3.19
WORKDIR /
COPY --from=builder /main ./
COPY ./config.env.example ./
COPY ./templates/ ./templates
ENTRYPOINT ["./main"]