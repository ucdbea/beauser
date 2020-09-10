FROM golang:1.15.1-alpine3.12 AS build-env
ENV CG0_ENABLED 0
RUN apk add --no-cache git

# RUN apk add --no-cache openssh-client git
# download public key for github.com
 #RUN mkdir -p -m 0600 ~/.ssh && ssh-keyscan github.com >> ~/.ssh/known_hosts
ADD . /go/src/beauser
#RUN --mount=type=ssh git clone https://github.com/ucdbea/beauser
# Grab the source code and add it to the workspace.


# Install revel and the revel CLI.
RUN go get  github.com/revel/revel
RUN go get  github.com/revel/cmd/revel

RUN git clone https://github.com/ucdbea/beauser app
RUN revel build -a app/beauser -m dev -v
# RUN  --mount=type=ssh revel build -a github.com/ucdbea/beauser  
# FROM alpine:3.8
# Use the revel CLI to start up our application.
# Open up the port where the app is running.
FROM alpine:edge
EXPOSE 9000
COPY --from=build-env /go/target /
COPY --from=build-env /go/src/beauser /opt/src
WORKDIR /




# WORKDIR /

# COPY --from=build-env . .
# ENTRYPOINT /run.sh
# CMD ["/run.sh"]
# EXPOSE 8080