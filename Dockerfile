########################
# Builder for artifact #
########################

FROM golang:1.11-alpine AS builder

# Set to your build args
ARG projectName
ARG projectOwner
ARG appName

RUN apk update && apk add --no-cache alpine-sdk openssh zip unzip

# Make sure you have github key into your known hosts for dep
RUN mkdir ~/.ssh && ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts

# Set the workdir to the project
WORKDIR /go/src/github.com/${projectOwner}/${projectName}

# Get the source in and build
COPY . .

# Install tools needed for the build
RUN go get -u github.com/golang/dep/cmd/dep && \
    dep ensure --vendor-only && \
    make build


#########################
# Build docker artifact #
#########################

FROM alpine:latest

ARG projectName
ARG projectOwner
ARG appName

ENV GIN_MODE="release"
ENV PORT=":3000"

RUN apk --no-cache add ca-certificates

WORKDIR /opt

COPY --from=builder /go/src/github.com/${projectOwner}/${projectName}/build/${appName}-linux-amd64 ${appnName}

EXPOSE 3000

CMD ["./microservice-template-linux-amd64"]

