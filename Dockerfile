################ DEBUG #######################
FROM golang:1.16.4-alpine3.13 as debug

# installing git
RUN apk update && apk upgrade && \
    apk add --no-cache git \
        dpkg \
        gcc \
        git \
        musl-dev

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# Get Go stuff
RUN go get github.com/go-delve/delve/cmd/dlv

WORKDIR /go/src/work
RUN pwd && ls -la /go/src/work
COPY ./ /go/src/work/

RUN go build -o app
### Run the Delve debugger ###
COPY ./deployment/dlv.sh /
RUN chmod +x /dlv.sh 
ENTRYPOINT [ "/dlv.sh"]


################# PROD ######################

FROM alpine:3.9 as prod
COPY --from=debug /go/src/work/app /
CMD ./app