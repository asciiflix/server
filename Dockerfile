################# PROD ######################
FROM golang:1.16.0-alpine AS builder
COPY . /server
WORKDIR /server
ENV GO111MODULE=on
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o /main .

FROM scratch
WORKDIR /
COPY --from=builder /main ./
ENTRYPOINT ["./main"]