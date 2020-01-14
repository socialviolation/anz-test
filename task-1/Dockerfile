# Builder
FROM golang:1.13-buster as build
WORKDIR /go/src/anz-test
ADD . /go/src/anz-test
RUN GOPROXY=https://proxy.golang.org go mod download
RUN go build -o /go/bin/anz-test

# Base
FROM gcr.io/distroless/base-debian10
COPY --from=build /go/bin/anz-test /
ARG APP_VERSION
ARG SHA
ENV APP_VERSION=${APP_VERSION} SHA=${SHA}
EXPOSE 8080
CMD ["/anz-test"]