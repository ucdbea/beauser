FROM golang:1.15.1-alpine3.12 AS build-env
RUN apk add --no-cache git
ADD . /go/src/beauser
# Grab the source code and add it to the workspace.


# Install revel and the revel CLI.
RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel

RUN revel build https://github.com/ucdbea/beauser /go/app
FROM alpine:3.8
# Use the revel CLI to start up our application.
# Open up the port where the app is running.
EXPOSE 9000

WORKDIR /

COPY --from=build-env /go /
ENTRYPOINT /run.sh