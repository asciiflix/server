################# PROD ######################

FROM golang:1.16.0-alpine AS builder
ARG VERSION=PROD
COPY . /server
WORKDIR /server
ENV GO111MODULE=on
RUN CGO_ENABLED=0 go build -ldflags "-X 'github.com/asciiflix/server/config.Version=$VERSION'" -o /main .

FROM scratch
WORKDIR /
COPY --from=builder /main ./
COPY ./config.env ./
ENTRYPOINT ["./main"]