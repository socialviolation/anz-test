# Builder
FROM golang:1.13-buster as build
WORKDIR /go/src/anz-test
ADD . /go/src/anz-test
RUN go mod vendor
RUN go build -o /go/bin/anz-test

# Base
FROM gcr.io/distroless/base-debian10
COPY --from=build /go/bin/anz-test /
EXPOSE 8080
CMD ["/anz-test"]