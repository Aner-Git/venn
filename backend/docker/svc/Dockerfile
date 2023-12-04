# syntax = docker/dockerfile:1.4
FROM golang:1.20 as local

RUN mkdir -p /go/src/github.com/meatmath.com
COPY . /go/src/github.com/meatmath.com/backend
WORKDIR /go/src/github.com/meatmath.com/backend

ARG service_name

# https://github.com/moby/buildkit/blob/master/frontend/dockerfile/docs/syntax.md#example-cache-go-packages
RUN --mount=type=cache,target=/root/.cache/go-build  go build -mod vendor -o /svc ./cmd/svc/${service_name}
RUN mkdir -p ./cmd/svc/${service_name}/resources
RUN cp -r ./cmd/svc/${service_name}/resources /resources
COPY ./docker/svc/resources.json /resources.json

CMD /svc

